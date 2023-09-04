package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewUploadVideoService new CheckUserService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	var err error
	// 查询视频是否存在
	_, err = db.CheckVideoExistById(req.VideoId)
	if err != nil {
		return err
	}
	// 参数出错
	if req.ActionType != 1 && req.ActionType != 2 {
		return errno.ParamErr
	}

	// 1 点赞
	if req.ActionType == 1 {
		return db.Add(s.ctx, req)
	} else {
		// 2 取消点赞
		return db.Delete(s.ctx, req)
	}

}
