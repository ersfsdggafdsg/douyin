package service

import (
	"context"
	"douyin/cmd/relation/pkg/manager"
	"douyin/cmd/relation/pkg/dal/redis"
	"gorm.io/gorm"
	"errors"
)

// 返回值(bool, error)
// (true/false, nil) 表示查询成功，
// (true/false, err) 表示服务器异常
func IsFollowing(m *manager.Manager, ctx context.Context, userId int64, fanId int64) (resp bool, err error) {
	resp, err = m.Rdb.IsFollowing(fanId, userId)
	switch err {
	case redis.Nil:
		_, err = m.Db.FollowInfo(userId, fanId)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp = false
		} else if err != nil {
			return false, err
		} else {
			resp = true
		}

		m.Rdb.FollowingSet(fanId, userId, resp)
		return resp, nil
	case nil:
		return resp, nil
	default:
		return false, err
	}

}
