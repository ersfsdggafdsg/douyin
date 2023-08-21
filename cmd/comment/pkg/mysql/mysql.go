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
	return config.DB.Save(&model.Comment {
		UserId: userId,
		VideoId: videoId,
		Content: content,
	}).Error
}

func CommentDel(commentId int64) error {
	return config.DB.Delete(&model.Comment{}, commentId).Error
}

func CommentList(userId int64, videoId int64) ([]*model.Comment, error) {
	list := make([]*model.Comment, 0)
	err := config.DB.Where(
		"user_id = ? AND video_id = ?", userId, videoId).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
