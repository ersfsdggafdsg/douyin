package manager

import (
	"douyin/cmd/favorite/pkg/dal/mysql"
	"douyin/cmd/favorite/pkg/dal/redis"
	"douyin/cmd/favorite/pkg/mq"
)
type Manager struct {
	Db mysql.FavoriteManager
	Rdb redis.RedisManager
	Mq mq.MessageQueueManager
}
