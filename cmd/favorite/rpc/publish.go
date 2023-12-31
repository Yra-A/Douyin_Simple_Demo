package rpc

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/middleware"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"time"

	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish/publishservice"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublishRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress}) // 服务发现
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
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
	publishClient = c
}

// GetVideoList 根据视频 id 列表获取视频列表【rpc 客户端】
func GetVideoList(ctx context.Context, req *publish.GetVideoListRequest) (*publish.GetVideoListResponse, error) {
	resp, err := publishClient.GetVideoList(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
