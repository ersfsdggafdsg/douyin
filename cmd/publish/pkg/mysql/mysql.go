package mysql

import (
	"douyin/cmd/publish/pkg/model"
	"douyin/shared/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type PublishManager struct {
	*gorm.DB
}

func NewManger(db *gorm.DB) (PublishManager) {
	return PublishManager{db}
}

// 创建一个视频记录
func (m *PublishManager)Create(info *model.VideoInfo) error {
	return m.Save(info).Error
}

// 更新视频的评论数
func (m *PublishManager)UpdateCommentCount(videoId int64, addCount int64) error {
	return m.Model(&model.VideoInfo{}).Scopes(utils.IdScope(videoId)).Update(
		"comment_count", gorm.Expr("comment_count + ?", addCount)).Error
}

// 更新视频的点赞数
func (m *PublishManager)UpdateFavoriteCount(videoId int64, addCount int64) error {
	return m.Model(&model.VideoInfo{}).Scopes(utils.IdScope(videoId)).Update(
		"favorite_count", gorm.Expr("favorite_count + ?", addCount)).Error
}

// 倒序查找创建时间小于t的最新的limit个视频
func (m *PublishManager)recent(t time.Time, limit int64) func(*gorm.DB) *gorm.DB {
	// 使用gorm的特性，排序、查找比指定时间早的表项，
	// 并且限制最终查到的数量。但是由于调用方要求使用的是函数，
	// 所以实际返回的也是函数。
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("created_at < ?", t).
			Order("created_at desc").Limit(int(limit))
	}
}

// 获取创建时间小于startTime的最新的limit个视频，如果它为0，就从当前开始
func (m *PublishManager)QueryRecentVideoInfos(startTime int64, limit int64) ([]*model.VideoInfo, error) {
	t := time.Unix(startTime / 1000, startTime % 1000 * 1000 * 1000)
	if startTime == 0 {
		t = time.Now()
	}
	infos := make([]*model.VideoInfo, 0)
	err := m.Scopes(m.recent(t, limit)).Find(&infos).Error
	return infos, err
}

// 获取用户的所有作品
func (m *PublishManager)QueryByAuthor(authorId int64) ([]*model.VideoInfo, error) {
	infos := make([]*model.VideoInfo, 0)
	err := m.Where("author_id = ? ", authorId,
		).Find(&infos).Error
	return infos, err
}

// 获得某个视频的信息
func (m *PublishManager)QueryById(videoId int64) (*model.VideoInfo, error) {
	info := new(model.VideoInfo)
	err := m.Scopes(utils.IdScope(videoId)).First(info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		return nil, err
	} else {
		return info, nil
	}
}

