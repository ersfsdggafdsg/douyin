package model

import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"gorm.io/gorm"
	"time"
)

type VideoInfo struct {
	rpc.VideoInfo
	DeleteAt gorm.DeletedAt
	CreateAt time.Time
	UpdateAt time.Time
}

func (VideoInfo)TableName() string {
	return "videos"
}

