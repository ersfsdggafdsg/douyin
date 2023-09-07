package model

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId int64
	VideoId int64
}

func (Favorite)TableName() string {
	return "favorites"
}
