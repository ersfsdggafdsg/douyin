package manager

import (
	"douyin/cmd/comment/pkg/dal/mysql"
	"douyin/cmd/comment/pkg/mq"
)

type Manager struct {
	Db mysql.CommentManager
	Mq mq.MessageQueueManager
}
