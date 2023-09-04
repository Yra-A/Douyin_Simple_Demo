package dal

import (
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/db"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/favorite/dal/redis"
)

// Init init dal
func Init() {
	db.Init()    // mysql
	redis.Init() // redis
}
