package redis

import (
	"fmt"

	"douyin/shared/initialize"
	redis "github.com/redis/go-redis/v9"
	"douyin/shared/utils/redismap"
	"golang.org/x/net/context"
)

var Nil = redis.Nil

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

func (m *RedisManager) getKey(fanId, userId int64) string {
	return fmt.Sprintf("follow:%d:%d", fanId, userId)
}

func (m *RedisManager) FollowingSet(fanId, userId int64, favorited bool) (error) {
	// TODO: 在值中存入该关系的id。这样删除时可以避免遍历数据库。
	return m.mp.Set(m.ctx, m.getKey(fanId, userId), favorited)
}

func (m *RedisManager) FollowingDel(fanId, userId int64) (error) {
	return m.mp.Del(m.ctx, m.getKey(fanId, userId))
}

// (*, error) 表示出错，error为redis.Nil的时候表示没有记录，其他都视为错误
// (*, nil) 才表示是否关注了
func (m RedisManager) IsFollowing(fanId, userId int64) (bool, error) {
	cmd, err := m.mp.Get(m.ctx, m.getKey(fanId, userId))
	if err != nil {
		return false, err
	}

	ret, err := cmd.Bool()
	return ret, err
}

