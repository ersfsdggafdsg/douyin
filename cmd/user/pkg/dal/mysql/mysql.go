package mysql

import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/cmd/user/pkg/dal/model"
	"douyin/shared/initialize"
	"gorm.io/gorm"
	"douyin/shared/utils"
)

type UserManager struct {
	*gorm.DB
}

func NewManager() (UserManager) {
	return UserManager{initialize.InitMysql(
		"douyin", "zhihao", "douyin", &model.UserInfo{})}
}

type DbTransaction struct {
	// 这一层封装的理由：
	// 1. DbTransaction和UserManager应该要有相同的方法
	// 2. 目前还没有使用池化技术，所以这么做省事一些。
	// 3. 以后用上了池化技术，大概也是先选择数据库链接，再操作。
	//    比如m.GetConn().UserAdd()这样的。
	UserManager
}

func (m *UserManager) RunTransaction(f func(tx *DbTransaction) error) error {
	// 传入参数是*DbTransaction，这是对UserManger的封装。
	// 做这一层封装的理由：
	// 1. DbTransaction和UserManager应该要有相同的方法
	// 2. 如果使用了池化技术，可以更方便的进行修改。
	//
	// 不使用Transaction这个名字的原因是
	// 1. 避免和gorm的那个Transaction重名
	// 2. 目前还没使用池化技术
	// 
	// 使用连接池后，这里可能是这样的了：
	// return m.GetConn().Transaction(func......)
	return m.Transaction(func(tx* gorm.DB) error {
		return f(&DbTransaction{
			UserManager{tx},
		})
	})
}

func (m *UserManager) UserAdd(userName, password string) (*model.UserInfo, error) {
	user := &model.UserInfo {
		UserInfo: rpc.UserInfo {
			Name: userName,
		},
		Password: password,
	}
	err := m.Create(user).Error
	return user, err
}

func (m *UserManager) UserList(userId int64) ([]*model.UserInfo, error) {
	infos := make([]*model.UserInfo, 0)
	err := m.Where("user_id = ?", userId).Find(&infos).Error
	return infos, err
}

// (*, gorm.RecordNotFound)，表示没有记录
// (*, 非nil)，当作服务器异常
// (对象, nil)，表示查询成功
func (m *UserManager) QueryByName(userName string) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := m.First(info, "name = ?", userName).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (m *UserManager) QueryById(id int64) (*model.UserInfo, error) {
	info := new(model.UserInfo)
	err := m.First(info, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

// 更新粉丝数量
func (m *UserManager) UpdateFollowerCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"follower_count", gorm.Expr("follower_count + ?", addCount)).Error
}

// 更新正在关注的人数
func (m *UserManager) UpdateFollowingCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"follow_count", gorm.Expr("follow_count + ?", addCount)).Error
}

// 更新点赞数
func (m *UserManager) UpdateFavoritingCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"favorite_count", gorm.Expr("favorite_count + ?", addCount)).Error
}

// 更新被(Be)点赞(Favorited)数
func (m *UserManager) UpdateBeFavoritedCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"total_favorited", gorm.Expr("total_favorited + ?", addCount)).Error
}


func (m *UserManager) UpdateWorkCount(id int64, addCount int64) error {
	return m.Model(&model.UserInfo{}).Scopes(utils.IdScope(id)).Update(
		"work_count", gorm.Expr("work_count + ?", addCount)).Error
}
