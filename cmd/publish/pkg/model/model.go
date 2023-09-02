package model

import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"gorm.io/gorm"
	"time"
)

type VideoInfo struct {
	// 这使用了嵌入，使得程序有类似于继承一般的用法
	rpc.VideoInfo
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (VideoInfo)TableName() string {
	return "videos"
}

