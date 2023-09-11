package mysql

import (
	"douyin/cmd/relation/pkg/dal/model"
	"douyin/shared/initialize"
	"gorm.io/gorm"
)

type RelationManager struct {
	*gorm.DB
}

func NewManager() (RelationManager) {
	return RelationManager{initialize.InitMysql(&model.Relation{})}
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
	RelationManager	
}

// 执行事务，参数是进行更新的函数。它执行成功返回nil，其他情况会自动回滚
func (m *RelationManager) RunTransaction(f func(tx *DbTransaction) error) error {
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
			RelationManager{tx},
		})
	})
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
