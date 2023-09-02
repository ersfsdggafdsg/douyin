package mysql

import (
	"douyin/cmd/favorite/pkg/model"

	"gorm.io/gorm"
)

type FavoriteManager struct {
	*gorm.DB
}

func NewManager(db *gorm.DB) (FavoriteManager) {
	return FavoriteManager{db}
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
