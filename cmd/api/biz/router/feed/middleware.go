// Code generated by hertz generator.

package feed

import (
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}
