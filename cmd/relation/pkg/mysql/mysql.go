package mysql

import (
	"douyin/cmd/relation/pkg/model"
	"gorm.io/gorm"
)

type RelationManager struct {
	*gorm.DB
}

func NewManager(db *gorm.DB) (RelationManager) {
	return RelationManager{db}
}

func (m *RelationManager)RelationAdd(fanId, userId int64) (*model.Relation, error) {
	r := &model.Relation {
		UserId: userId,
		FanId: fanId,
	}
	return r, m.Save(r).Error
}

func (m *RelationManager)RelationDel(fanId, userId int64) (*model.Relation, error) {
	info := new(model.Relation)
	result := m.Where("user_id = ? AND fan_id = ?",
		userId, fanId).Delete(info)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	} else {
		return info, nil
	}
}

// userId的粉丝
func (m *RelationManager)FansList(userId int64) ([]int64, error) {
	var ids []int64
	err := m.Model(&model.Relation{}).
		Where("user_id = ?", userId).
		Pluck("fan_id", &ids).Error
	return ids, err
}

// 正在关注的人，即userId作为粉丝
func (m *RelationManager)FollowList(userId int64) ([]int64, error) {
	var ids []int64
	err := m.Model(&model.Relation{}).
		Where("fan_id = ?", userId).
		Pluck("user_id", &ids).Error
	return ids, err
}

func (m *RelationManager)FriendList(userId int64) ([]int64, error) {
	// 实现是一样的，多的一步是获取最新的聊天消息
	// 但是不是放在这实现的
	return m.FansList(userId)
}

/* 查找某个关系的信息
 * 如果找到了，返回Follow对象，nil
 * 如果找不到，返回nil，gorm.ErrRecordNotFound
 * 其他情况，nil和error
 * 总之，找到了返回对象和nil，其他情况是nil和error
 */
func (m *RelationManager)FollowInfo(userId, fanId int64) (*model.Relation, error) {
	r := new(model.Relation)
	err := m.First(r, "user_id = ? AND fan_id = ?", fanId, userId).Error
	if err != nil {
		r = nil
	}
	return r, err
}
