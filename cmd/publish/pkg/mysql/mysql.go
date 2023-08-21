package mysql

import (
	"douyin/cmd/publish/pkg/model"
	"douyin/shared/config"
	"douyin/shared/initialize"
	"douyin/shared/rpc/kitex_gen/rpc"
	"time"

	"gorm.io/gorm"
)

func init() {
	initialize.InitMysql(&model.VideoInfo{})
}

// 创建一个视频记录
func Create(info *model.VideoInfo) error {
	return config.DB.Create(info).Error
}

// 更新视频的评论数
func UpdateCommentCount(videoId int64, newCount int64) error {
	return config.DB.Model(&model.VideoInfo{}).Where(
			"id = ?", videoId,
		).Update(
			"comment_count", newCount,
		).Error
}

// 更新视频的点赞数
func UpdateFavoriteCount(videoId int64, newCount int64) error {
	return config.DB.Model(&model.VideoInfo{}).Where(
			"id = ?", videoId,
		).Update(
			"favorite_count", newCount,
		).Error
}

// 倒序查找创建时间小于t的最新的limit个视频
func recent(t time.Time, limit int64) func(*gorm.DB) *gorm.DB {
	// 使用gorm的特性，排序、查找比指定时间早的表项，
	// 并且限制最终查到的数量。但是由于调用方要求使用的是函数，
	// 所以实际返回的也是函数。
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("create_at <= t", t).
			Order("create_at desc").Limit(int(limit))
	}
}

// 获取创建时间小于startTime的最新的limit个视频
func QueryRecentVideoInfos(startTime int64, limit int64) ([]*model.VideoInfo, error) {
	t := time.Unix(startTime, 0)
	infos := make([]*model.VideoInfo, 0)
	err := config.DB.Scopes(recent(t, limit)).Find(&infos).Error
	return infos, err
}

// 获取用户的所有作品
func GetUserWorks(authorId int64) ([]*model.VideoInfo, error) {
	infos := make([]*model.VideoInfo, 0)
	err := config.DB.Where(
			"author_id = ? ", authorId,
		).Find(&model.VideoInfo{}).Error
	return infos, err
}

// 获得某个视频的信息
func GetVideoInfo(videoId int64) (*rpc.VideoInfo, error) {
	infos := make([]*model.VideoInfo, 0)
	result := config.DB.Where("id = ?", videoId).Find(infos, "limit = 1")
	if result.Error != nil {
		return nil, result.Error
	}
	return &infos[0].VideoInfo, nil
}
