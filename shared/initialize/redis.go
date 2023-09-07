package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	redis "github.com/redis/go-redis/v9"
)


func InitRedis() (client *redis.Client) {
	rdb := redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	if rdb == nil {
		klog.Fatal("Can't create redis client")
	}

	return rdb
}

