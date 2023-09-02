package model

import (
	"douyin/shared/rpc/kitex_gen/common"
	"gorm.io/gorm"
)

type Message struct {
	common.Message
	DeleteAt gorm.DeletedAt
}

func (Message) TableName() string {
	return "messages"
}
