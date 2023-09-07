package mysql

import (
	"douyin/cmd/favorite/pkg/dal/model"
	"douyin/shared/initialize"

	"gorm.io/gorm"
)

type FavoriteManager struct {
	*gorm.DB
}

func NewManager() (FavoriteManager) {
	return FavoriteManager{initialize.InitMysql(
			"douyin", "zhihao", "douyin", &model.Favorite{},
		)}
}

type DbTransaction struct {
	// 为什么这么做？
	// 这么做只是为了给是否使用只读事务，留下选择的余地
	// 
	// 关于是否要使用只读事务，网上有很多说法。
	// 有人说不要，有人说要。
	// 说不要的认为这降低了性能，说要的认为这可以确保操作的原子性，
	// 并且可以设置锁超时和低隔离级别来提升性能。
	// 
	// 这里给了这么做的理由（仅我个人观点）：
	// 1. 写入请求肯定是要使用事务的
	// 2. 一些未来可能需要加入写入的读取，需要事务。
	// 3. 大概率不需要加上写入的地方则不需要。
	// 4. gorm的事务支持大多数数据库操作
	FavoriteManager
}

// 执行事务，参数是进行更新的函数。它执行成功返回nil，其他情况会自动回滚
func (m *FavoriteManager) RunTransaction(f func(tx *DbTransaction) error) error {
	// 传入参数是*DbTransaction，这是对FavoriteManager的封装。
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
			FavoriteManager{tx},
		})
	})
}

// 进行点赞
// (nil, error)表示失败，error为gorm.ErrDuplicatedKey表示重复点赞
// (*model.Favorite, nil)表示成功
func (m *FavoriteManager) FavoriteAdd(userId, videoId int64) (*model.Favorite, error) {
	info := &model.Favorite {
		UserId: userId,
		VideoId: videoId,
	}
	err := m.Save(info).Error
	return info, err
}

// 取消点赞
// (nil, error)表示失败
//     error为gorm.ErrRecordNotFound表示在没有点赞的情况下取消点赞
// (*model.Favorite, nil)表示成功
func (m *FavoriteManager) FavoriteDel(userId, videoId int64) (*model.Favorite, error) {
	info := new(model.Favorite)
	result := m.Delete(info, "user_id = ? AND video_id = ?", userId, videoId)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return info, nil
}

// 获取点赞列表，err为nil表示成功
func (m *FavoriteManager) FavoriteList(userId int64) ([]*model.Favorite, error) {
	infos := make([]*model.Favorite, 0)
	err := m.Where("user_id = ?", userId).Find(&infos).Error
	return infos, err
}

// 获取点赞信息
//     (nil, gorm.ErrRecordNotFound)表示没有点赞
//     (nil, error)表示查询出错
// 唯有(*model.Favorite{}, nil)表示找到了信息
func (m *FavoriteManager) FavoriteInfo(userId, videoId int64) (*model.Favorite, error) {
	f := new(model.Favorite)
	err := m.Where("user_id = ? AND video_id = ?",
		userId, videoId).First(f).Error
	if err != nil {
		return nil, err
	} else {
		return f, nil
	}
}
