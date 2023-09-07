package redis

import (
	"fmt"

	"douyin/shared/initialize"
	redis "github.com/redis/go-redis/v9"
	"douyin/shared/utils/redismap"
	"golang.org/x/net/context"
)

type RedisManager struct {
	cli *redis.Client
	mp   *redismap.ExpireMap
	ctx context.Context
}

func NewManager() (RedisManager) {
	db := initialize.InitRedis()
	return RedisManager{
		cli: db,
		mp : redismap.NewExpireMap(db),
		ctx: context.Background(),
	}
}

func (m *RedisManager) getKey(userId, videoId int64) string {
	return fmt.Sprintf("favorite:%d:%d", userId, videoId)
}

func (m *RedisManager) FavoriteSet(userId, videoId int64, favorited bool) (error) {
	return m.mp.Set(m.ctx, m.getKey(userId, videoId), favorited)
}

func (m *RedisManager) FavoriteDel(userId, videoId int64) (error) {
	return m.mp.Del(m.ctx, m.getKey(userId, videoId))
}

// (*, error) 表示出错，error为redis.Nil的时候表示没有记录，其他都视为错误
// (*, nil) 才表示是否关注了
func (m RedisManager) IsFavorited(userId, videoId int64) (bool, error) {
	cmd, err := m.mp.Get(m.ctx, m.getKey(userId, videoId))
	if err != nil {
		return false, err
	}

	ret, err := cmd.Bool()
	return ret, err
}
