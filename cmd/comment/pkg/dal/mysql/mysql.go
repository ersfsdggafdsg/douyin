package mysql

import (
	"douyin/cmd/comment/pkg/dal/model"
	"douyin/shared/initialize"

	"gorm.io/gorm"
)

type CommentManager struct {
	*gorm.DB
}

func NewManager() (CommentManager) {
	return CommentManager{initialize.InitMysql(
		"douyin", "zhihao", "douyin", &model.Comment{})}
}

type DbTransaction struct {
	// 这一层封装的理由：
	// 1. DbTransaction和CommentManager应该要有相同的方法
	// 2. 目前还没有使用池化技术，所以这么做省事一些。
	// 3. 以后用上了池化技术，大概也是先选择数据库链接，再操作。
	//    比如m.GetConn().CommentAdd()这样的。
	CommentManager
}

// 执行事务，参数是进行更新的函数。它执行成功返回nil，其他情况会自动回滚
func (m *CommentManager) RunTransaction(f func(tx *DbTransaction) error) error {
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
			CommentManager{tx},
		})
	})
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

