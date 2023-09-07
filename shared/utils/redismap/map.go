package redismap

import (
	"context"
	"time"

	"math/rand"
	redis "github.com/redis/go-redis/v9"
)

// WARN: Github上有回答：
// The client is thread safe, but pipelines, multi blocks, etc are not.

type ExpireMap struct {
	rdb *redis.Client
}

func randTime() (time.Duration) {
	// 返回的是一个小时+随机秒数
	return time.Hour + time.Second * time.Duration(rand.Intn(60))
}

// 使用redis来实现，它自带过期时间。
func NewExpireMap(cli *redis.Client) *ExpireMap {
	return &ExpireMap {
		rdb: cli,
	}
}

func (m *ExpireMap)Set(ctx context.Context, key string, value interface{}) error {
	return m.rdb.Set(ctx, key, value, randTime()).Err()
}

func (m *ExpireMap)Del(ctx context.Context, key string) error {
	return m.rdb.Del(ctx, key).Err()
}

func (m *ExpireMap)Exists(ctx context.Context, key string) (bool, error) {
	result, err := m.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, nil
	}

	return true, nil
}

func (m *ExpireMap)Get(ctx context.Context, key string) (*redis.StringCmd, error) {
	// 是否需要每次读取的时候更新过期时间？
	// 我觉得不能自动更新过期时间。理由是，
	// 缓存和数据库很难保持一致，如果还自动更新过期时间，那从不一致到一致的周期更长。
	// 在这个项目里，用户、视频表里有点赞数量等信息，他们需要走消息队列才能够更新。
	// 如果消息发布成功而数据库事务提交失败，那么数据就已经不一致了，
	// 所以自动过期时间还是短一些好。
	// 所以在表中去掉这些数量记录，是后续的一个更改。
	cmd := m.rdb.Get(ctx, key)
	err := cmd.Err()
	return cmd, err
}

