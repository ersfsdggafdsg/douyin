package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	redis "github.com/redis/go-redis/v9"
)


func InitRedis() (client *redis.Client) {
	rdb := redis.NewClient(&redis.Options {
		Addr: Config.GetString("redis_addr"),
		Password: Config.GetString("redis_password"),
		DB: Config.GetInt("redis_db_no"),
	})

	if rdb == nil {
		klog.Fatal("Can't create redis client")
	}

	return rdb
}

