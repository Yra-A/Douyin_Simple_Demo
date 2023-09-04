package redis

import (
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress,
		Password: constants.RedisPassword, // no password set
		DB:       constants.DBIndex,       // use default DB
	})

}
