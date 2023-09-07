package service

import (
	"context"
	"douyin/cmd/favorite/pkg/manager"
	"errors"

	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)
func IsFavorited(m *manager.Manager, ctx context.Context, userId, videoId int64) (resp bool, err error) {
	favorited, err := m.Rdb.IsFavorited(userId, videoId)
	if errors.Is(err, redis.Nil) {
		// 键不存在，往下查询数据库
	} else if err != nil {
		// 出现错误，返回
		// 为什么这么设计？
		// redis崩溃，很可能只是其中一个服务器崩溃了，整个集群还能用。
		// 大不了叫用户重试就是了。
		return false, err
	} else {
		return favorited, nil
	}

	_, err = m.Db.FavoriteInfo(userId, videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		m.Rdb.FavoriteSet(userId, videoId, false)
		return false, nil
	} else if err != nil {
		return false, err
	}

	m.Rdb.FavoriteSet(userId, videoId, true)
	return true, nil
}
