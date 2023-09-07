package redismap

import (
	"fmt"
	"os"
	"testing"

	redis "github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestMap(t *testing.T) {
	mp := ExpireMap{redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:		  0,  // 默认DB 0
	})}
	// 设置键
	mp.Add(context.Background(), "v", 10)

	keys := []string {"v", "not"}
	// 获得键
	for _, v := range keys {
		str, err := mp.Get(context.Background(), v)
		fmt.Println(v, str, err)
	}
}
