package mysql

import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/cmd/user/pkg/model"
	"gorm.io/gorm"
	"douyin/shared/utils"
)

type UserManager struct {
	*gorm.DB
}

func NewManager(db *gorm.DB) (UserManager) {
	return UserManager{db}
}

func (m *UserManager)UserAdd(userName, password string) (*model.UserInfo, error) {
	user := &model.UserInfo {
		UserInfo: rpc.UserInfo {
			Name: userName,
		},
		Password: password,
	}
	err := m.Create(user).Error
	return user, err
}

func (m *UserManager)UserList(userId int64) ([]*model.UserInfo, error) {
	infos := make([]*model.UserInfo, 0)
	err := m.Where("user_id = ?", userId).Find(&infos).Error
	return infos, err
}

func (m *UserManager)QueryByName(userName string) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := m.First(info, "name = ?", userName).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (m *UserManager)QueryById(id int64) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := m.First(info, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

// 更新粉丝数量
func (m *UserManager)UpdateFollowerCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"follower_count", gorm.Expr("follower_count + ?", addCount)).Error
}

// 更新正在关注的人数
func (m *UserManager)UpdateFollowingCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"follow_count", gorm.Expr("follow_count + ?", addCount)).Error
}

// 更新点赞数
func (m *UserManager)UpdateFavoritedCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"favorite_count", gorm.Expr("favorite_count + ?", addCount)).Error
}

func (m *UserManager)UpdateWorkCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"work_count", gorm.Expr("work_count + ?", addCount)).Error
}
