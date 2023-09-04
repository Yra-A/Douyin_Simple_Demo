package service

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type FavoriteCountByUserIdService struct {
	ctx context.Context
}

// NewFavoriteCountByUserIdService new FavoriteCountByUserIdService
func NewFavoriteCountByUserIdService(ctx context.Context) *FavoriteCountByUserIdService {
	return &FavoriteCountByUserIdService{
		ctx: ctx,
	}
}

func (s *FavoriteCountByUserIdService) FavoriteCountByUserId(req *favorite.FavoriteCountByUserIDRequest) (int64, error) {
	//拉取
	cnt, err := db.GetFavoriteCountByUserID(req.UserId)
	return cnt, err
}
