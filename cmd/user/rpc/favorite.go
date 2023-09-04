package rpc

import (
	"context"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite/favoriteservice"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	resp, err := favoriteClient.FavoriteCount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	resp, err := favoriteClient.IsFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FavoriteCountByUserId(ctx context.Context, req *favorite.FavoriteCountByUserIDRequest) (*favorite.FavoriteCountByUserIDResponse, error) {
	resp, err := favoriteClient.FavoriteCountByUserID(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TotalFavoritedByAuthorId(ctx context.Context, req *favorite.TotalFavoritedByAuthorIDRequest) (*favorite.TotalFavoritedByAuthorIDResponse, error) {
	resp, err := favoriteClient.TotalFavoritedByAuthorID(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
