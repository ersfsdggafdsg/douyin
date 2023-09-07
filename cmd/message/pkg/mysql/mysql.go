package mysql

import (
	"douyin/cmd/message/pkg/model"
	"douyin/shared/utils"
	"time"
	"douyin/shared/rpc/kitex_gen/common"

	"gorm.io/gorm"
)

type MessageManager struct {
	*gorm.DB
}

func NewManager(db *gorm.DB) MessageManager {
	return MessageManager{db}
}

func (m *MessageManager) LatestMessage(senderId, receiverId int64) (resp *model.Message, err error) {
	resp = new(model.Message)
	// 这样避免查询两次
	err = m.Order("create_time desc").First(&resp,
		"(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
		senderId, receiverId, receiverId, senderId).Error
	return resp, err
}

func (m *MessageManager) recent(preMsgTime int64) func(*gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		// FIXME: 可能实际使用的不是秒
		return db.Where("create_time > ?", utils.Time2Str(time.Unix(preMsgTime / 1000, preMsgTime % 1000 * 1000 * 1000)))
	}
}
func (m *MessageManager) MessageList(fromUserId, toUserId, preMsgTime int64) ([]*model.Message, error) {
	infos := make([]*model.Message, 0)
	err := m.Scopes(m.recent(preMsgTime)).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Find(&infos).Error
	return infos, err
}

func (m *MessageManager) MessageAdd(fromUserId , toUserId int64, content string) (*model.Message, error) {
	info := &model.Message {
			Message: common.Message {
				FromUserId: fromUserId,
				ToUserId  : toUserId,
				Content   : content,
				CreateTime: utils.Now2Str(),
			},
		}
	err := m.Save(info).Error
	return info, err
}
