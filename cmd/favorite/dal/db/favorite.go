package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/redis"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"gorm.io/gorm"
)

var rdFav redis.Favorite

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (u *Favorite) TableName() string {
	return constants.FavoriteTableName
}

type Video struct {
	gorm.Model
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

type User struct {
	gorm.Model
	ID              int64  `gorm:"primaryKey";json:"id"`
	UserName        string `gorm:"type:varchar(255)"json:"user_name"`
	Password        string `gorm:"type:varchar(255)"json:"password"`
	Avatar          string `gorm:"type:varchar(255)"json:"avatar"`           // 用户头像 URL
	BackgroundImage string `gorm:"type:varchar(255)"json:"background_image"` // 用户背景图 URL
	Signature       string `gorm:"type:varchar(255)"json:"signature"`        // 用户个性签名
}

func (Video) TableName() string {
	return constants.VideosTableName
}

func Delete(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	// 更新数据库
	err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	// 更新缓存
	if rdFav.CheckLiked(req.VideoId) {
		rdFav.DelLiked(req.UserId, req.VideoId)
	}
	if rdFav.CheckLike(req.UserId) {
		rdFav.DelLike(req.UserId, req.VideoId)
	}
	return nil
}

// Add favorite
func Add(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	fav := &Favorite{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	}
	// 更新数据库
	if err := DB.WithContext(ctx).Create(fav).Error; err != nil {
		return err
	}
	// 更新缓存
	if rdFav.CheckLiked(fav.VideoId) {
		rdFav.AddLiked(fav.UserId, fav.VideoId)
	}
	if rdFav.CheckLike(fav.UserId) {
		rdFav.AddLike(fav.UserId, fav.VideoId)
	}
	return nil
}

func GetFavoriteIdList(ctx context.Context, user_id int64) ([]int64, error) {
	if rdFav.CheckLike(user_id) {
		fmt.Println("通过 redis 缓存获取用户点赞视频列表成功！")
		return rdFav.GetLike(user_id), nil
	}
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", user_id).Find(&res).Error; err != nil {
		return nil, err
	}

	var video_id []int64
	for _, users := range res {
		video_id = append(video_id, users.VideoId)
	}
	return video_id, nil
}

func GetVideoListByVideoIDList(ctx context.Context, video_id_list []int64) ([]*Video, error) {
	var video_list []*Video
	var err error
	for _, item := range video_id_list {
		var video *Video
		err = DB.WithContext(ctx).Where("id = ?", item).Find(&video).Error
		if err != nil {
			return video_list, err
		}
		video_list = append(video_list, video)
	}

	return video_list, err
}

// QueryFavoriteCount favorite_count  video_id how many people like
func QueryFavoriteCount(ctx context.Context, video_id int64) (int64, error) {
	var favorite_count int64
	var err error
	if rdFav.CheckLiked(video_id) {
		fmt.Println("通过 redis 缓存获取视频获赞数量成功！")
		favorite_count, err = rdFav.CountLiked(video_id)
		if err != nil {
			return 0, err
		}
		return favorite_count, nil
	}
	favorite_count = 0
	likes := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", video_id).Find(&likes).Error; err != nil {
		return favorite_count, err
	}
	// 异步更新缓存
	go func(likes []*Favorite, video int64) {
		for _, like := range likes {
			rdFav.AddLiked(like.UserId, video)
			rdFav.AddLike(like.UserId, video)
		}
	}(likes, video_id)
	favorite_count = int64(len(likes))
	return favorite_count, nil
}

// QueryIsFavorite user_id like video_id
func QueryIsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (bool, error) {
	if rdFav.CheckLiked(req.VideoId) {
		fmt.Println("通过 redis 缓存获取用户是否点赞视频成功！")
		return rdFav.ExistLiked(req.UserId, req.VideoId), nil
	}
	if rdFav.CheckLike(req.UserId) {
		fmt.Println("通过 redis 缓存获取用户是否点赞视频成功！")
		return rdFav.ExistLike(req.UserId, req.VideoId), nil
	}
	res := make([]*Favorite, 0)
	DB.WithContext(ctx).Where("user_id = ? and video_id = ?", req.UserId, req.VideoId).Find(&res)
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

// GetFavoriteCountByUserID get the num of the video liked by user
func GetFavoriteCountByUserID(user_id int64) (int64, error) {
	if rdFav.CheckLike(user_id) {
		return rdFav.CountLike(user_id)
	}
	var likes []*Favorite
	if err := DB.Model(&Favorite{}).Where("user_id = ?", user_id).Find(&likes).Error; err != nil {
		return 0, err
	}

	// update redis asynchronously
	go func(user int64, likes []*Favorite) {
		for _, like := range likes {
			rdFav.AddLiked(user, like.VideoId)
			rdFav.AddLike(user, like.VideoId)
		}
	}(user_id, likes)

	return int64(len(likes)), nil
}

// QueryTotalFavoritedByAuthorID 获取该作者的获赞数
func QueryTotalFavoritedByAuthorID(user_id int64) (int64, error) {
	var count int64
	videos := make([]*Video, 0)
	if err := DB.Select("id").Where("author_id = ?", user_id).Find(&videos).Error; err != nil {
		return 0, err
	}

	for _, v := range videos {
		cnt, err := QueryFavoriteCount(context.Background(), v.ID)
		if err != nil {
			return 0, err
		}
		count += cnt
	}
	return count, nil
}

// CheckVideoExistById query if video exist
func CheckVideoExistById(video_id int64) (bool, error) {
	var video Video
	if err := DB.Where("id = ?", video_id).Find(&video).Error; err != nil {
		return false, err
	}
	if video == (Video{}) {
		return false, nil
	}
	return true, nil
}

// CheckUserExistById find if user exists
func CheckUserExistById(user_id int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
