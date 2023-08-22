package mysql

import (
	"douyin/cmd/comment/pkg/model"
	"douyin/shared/config"
	"douyin/shared/initialize"
)

func init() {
	initialize.InitMysql(&model.Comment{})
}

func CommentAdd(userId int64, videoId int64, content string) error {
	// 这个语句，如果userId(主键)不在表中，创建一个，否则更新。
	return config.DB.Save(&model.Comment {
		UserId: userId,
		VideoId: videoId,
		Content: content,
	}).Error
}

func CommentDel(commentId int64) error {
	return config.DB.Delete(&model.Comment{}, commentId).Error
}

func CommentList(videoId int64) ([]*model.Comment, error) {
	list := make([]*model.Comment, 0)
	// 这个语句相当于 SELECT * FROM comments where AND video_id = videoId
	// 查询的结果放在list中
	err := config.DB.Where(
		"video_id = ?", videoId).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
