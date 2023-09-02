package redisset
import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type Set struct {
	rdb *redis.Client
	name string
}

func NewSet(setName string, cli *redis.Client) *Set {
	return &Set {
		rdb: cli,
		name: setName,
	}
}

func (set *Set)Add(ctx context.Context, key string) error {
	return set.rdb.SAdd(ctx, set.name, key).Err()
}

func (set *Set)Del(ctx context.Context, key string) error {
	return set.rdb.SRem(ctx, set.name, key).Err()
}

func (set *Set)Exists(ctx context.Context, key string) (bool, error) {
	return set.rdb.SIsMember(ctx, set.name, key).Result()
}

func (set *Set)Size(ctx context.Context) (int64, error) {
	return set.rdb.SCard(ctx, set.name).Result()
}

