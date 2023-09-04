package service

import (
	"context"
	"fmt"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/feed/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/feed/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/feed"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(req *feed.FeedRequest) ([]*feed.Video, int64, error) {
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime/1000, 0)
	}
	fmt.Printf("视频的最新投稿时间戳为 %v\n", lastTime)

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return nil, 0, err
	}

	videos := make([]*feed.Video, 0, constants.MaxFeedCount)
	var nextTime int64
	if len(dbVideos) > 0 {
		nextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}

	for _, v := range dbVideos {
		favoriteCount, commentCount, isFavorite := s.getOtherVideoInfo(v.ID, req.UserId)
		author := s.getAuthorByUserId(v.AuthorID)
		nv := &feed.Video{
			Id:            v.ID,
			Author:        author,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			Title:         v.Title,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
		}
		videos = append(videos, nv)
	}

	return videos, nextTime, nil
}

// TODO: 实现获取视频相关信息 (并发获取?)
func (s *FeedService) getOtherVideoInfo(videoID int64, myId int64) (int64, int64, bool) {
	var favoriteCount int64
	var commentCount int64
	var isFavorite bool
	var wg sync.WaitGroup
	wg.Add(2)

	// 获取 favorite count
	go func() {
		defer wg.Done()
		resp, err := rpc.FavoriteCount(s.ctx, &favorite.FavoriteCountRequest{VideoId: videoID})
		favoriteCount = resp.FavoriteCount
		if err != nil {
			klog.CtxInfof(s.ctx, "获取 favorite count 出现错误："+err.Error())
		}
	}()

	// TODO 获取 comment count

	// 获取 is favorite
	go func() {
		defer wg.Done()
		resp, err := rpc.IsFavorite(s.ctx, &favorite.IsFavoriteRequest{UserId: myId, VideoId: videoID})
		isFavorite = resp.IsFavorite
		if err != nil {
			klog.CtxInfof(s.ctx, "获取 is favorite 出现错误："+err.Error())
		}
	}()

	wg.Wait()

	return favoriteCount, commentCount, isFavorite
}

func (s *FeedService) getAuthorByUserId(UserId int64) *feed.User {
	u, err := rpc.GetUser(s.ctx, &user.UserInfoRequest{
		UserId: UserId,
	})
	if err != nil {
		klog.CtxInfof(s.ctx, "获取视频的 Author 出现错误："+err.Error())
		return nil
	}
	return &feed.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}
