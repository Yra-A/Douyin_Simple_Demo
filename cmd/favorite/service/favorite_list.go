package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/redis"
	"log"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
)

var rdFav redis.Favorite

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList favorite list
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	_, err := db.CheckUserExistById(req.UserId)
	if err != nil {
		return nil, err
	}

	//本用户id + video_id[]  获取 video_list
	video_ids, _ := db.GetFavoriteIdList(s.ctx, req.UserId)
	if len(video_ids) == 0 {
		log.Println("FavoriteList : video_ids is blank")
	}

	temp, err := rpc.GetVideoList(s.ctx, &publish.GetVideoListRequest{UserId: req.UserId, VideoIds: video_ids})
	if err != nil {
		return nil, err
	}
	var resp []*favorite.Video
	for _, a := range temp.VideoList {
		b := &favorite.Video{Id: a.Id, Author: (*favorite.User)(a.Author), PlayUrl: a.PlayUrl, CoverUrl: a.CoverUrl, FavoriteCount: a.FavoriteCount, IsFavorite: a.IsFavorite, CommentCount: a.CommentCount, Title: a.Title}
		resp = append(resp, b)
	}

	return resp, err
}
