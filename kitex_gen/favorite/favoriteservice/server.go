// Code generated by Kitex v0.7.0. DO NOT EDIT.
package favoriteservice

import (
	favorite "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler favorite.FavoriteService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
