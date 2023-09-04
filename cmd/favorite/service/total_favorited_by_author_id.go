package service

import (
	"context"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
)

type TotalFavoritedByAuthorIdService struct {
	ctx context.Context
}

// NewTotalFavoritedByAuthorIdService new TotalFavoritedByAuthorIdService
func NewTotalFavoritedByAuthorIdService(ctx context.Context) *TotalFavoritedByAuthorIdService {
	return &TotalFavoritedByAuthorIdService{
		ctx: ctx,
	}
}

func (s *TotalFavoritedByAuthorIdService) TotalFavoritedByAuthorId(req *favorite.TotalFavoritedByAuthorIDRequest) (int64, error) {
	//拉取
	cnt, err := db.QueryTotalFavoritedByAuthorID(req.AuthorId)
	return cnt, err
}
