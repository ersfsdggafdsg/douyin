package service

import (
	"context"
	"douyin/cmd/user/pkg/dal/mysql"
	"douyin/cmd/user/pkg/manager"
)

func UpdateWorkCount(m *manager.Manager, ctx context.Context, userId int64, addCount int64) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		return tx.UpdateWorkCount(userId, addCount)
	})
}
