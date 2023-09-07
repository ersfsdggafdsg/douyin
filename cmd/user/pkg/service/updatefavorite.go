package service

import (
	"context"
	"douyin/cmd/user/pkg/dal/mysql"
	"douyin/cmd/user/pkg/manager"

	"github.com/cloudwego/kitex/pkg/klog"
)

// UpdateFavoritedCount implements the UserServiceImpl interface.
func UpdateFavoriteCount(m *manager.Manager, ctx context.Context, authorId, userId, addCount int64) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		err = tx.UpdateBeFavoritedCount(authorId, addCount)
		if err != nil {
			klog.Error("update favorite count failed:", err)
			return err
		}

		err = tx.UpdateFavoritingCount(userId, addCount)
		if err != nil {
			klog.Error("update favorite count failed:", err)
			return err
		}

		return nil
	})
}
