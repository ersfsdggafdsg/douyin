package model

import (
	"gorm.io/gorm"
)

// ToUserId还是直接叫粉丝来的愉快
type Relation struct {
	gorm.Model
	UserId int64
	FanId int64
}

func (Relation)TableName() string {
	return "relations"
}
