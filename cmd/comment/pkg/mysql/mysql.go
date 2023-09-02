package mysql

import (
	"douyin/cmd/comment/pkg/model"

	"gorm.io/gorm"
)

type CommentManager struct {
	*gorm.DB
}

func NewManager(db *gorm.DB) (CommentManager) {
	return CommentManager{db}
}

func (m *CommentManager) CommentAdd(userId int64, videoId int64, content string) (ret *model.Comment, err error) {
	// 这个语句，如果userId(主键)不在表中，创建一个，否则更新。
	ret = &model.Comment {
			UserId: userId,
			VideoId: videoId,
			Content: content,
		}
	err = m.Save(&ret).Error
	return
}

func (m *CommentManager) CommentDel(commentId int64) (*model.Comment, error) {
	comment := new(model.Comment)
	result :=  m.Delete(comment, commentId)

	// 这里是Gorm的一个决策，删除不存在的记录是不会报错的
	if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return comment, nil
}

func (m *CommentManager) CommentList(videoId int64) ([]*model.Comment, error) {
	list := make([]*model.Comment, 0)
	// 这个语句相当于 SELECT * FROM comments where AND video_id = videoId
	// 查询的结果放在list中
	err := m.Where(
		"video_id = ?", videoId).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

