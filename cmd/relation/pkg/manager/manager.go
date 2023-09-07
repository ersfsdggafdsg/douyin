package manager

import (
	"douyin/cmd/relation/pkg/mq"
	"douyin/cmd/relation/pkg/dal/mysql"
	"douyin/cmd/relation/pkg/dal/redis"
)

type Manager struct {
	Db mysql.RelationManager
	Rdb redis.RedisManager
	Mq mq.MessageQueueManager
}
