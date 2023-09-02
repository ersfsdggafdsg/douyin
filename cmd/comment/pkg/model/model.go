package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	// 这里和publish不同，原因是，实际上要存储的是用户ID而不是用户信息
	gorm.Model
	UserId int64
	VideoId int64
	Content string
}

func (Comment)TableName() string {
	return "comments"
}

