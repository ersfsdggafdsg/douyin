package service

import (
	"context"
	"douyin/cmd/user/pkg/dal/mysql"
	"douyin/cmd/user/pkg/manager"

	"github.com/cloudwego/kitex/pkg/klog"
)

// UpdateFollowdCount implements the UserServiceImpl interface.
func UpdateFollowCount(m *manager.Manager, ctx context.Context, userId, fanId, addCount int64) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		// 更新粉丝关注的人的数量
		err = tx.UpdateFollowingCount(fanId, addCount)
		if err != nil {
			klog.Error("update following count failed:", err)
			return err
		}

		// 更新作者粉丝数量
		err = tx.UpdateFollowerCount(userId, addCount)
		if err != nil {
			klog.Error("update follower count failed:", err)
			return err
		}

		return nil
	})
}
