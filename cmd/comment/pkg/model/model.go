package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId int64
	VideoId int64
	Content string
}

func (Comment)TableName() string {
	return "comments"
}

