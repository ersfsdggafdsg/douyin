package service

import (
	"context"
	"douyin/cmd/favorite/pkg/manager"
	"douyin/cmd/favorite/pkg/dal/mysql"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
	"errors"
	"gorm.io/gorm"
)

func FavoriteDel(m *manager.Manager, ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		// Transaction在返回nil的时候会自动提交，否则回滚
		err = favoriteDel(tx, ctx, req, resp)
		if err != nil {
			return err
		}
		
		// 这个顺序的理由也是数据库的数据更加重要，缓存大不了等他过期。
		err = updateRemote(ctx, &m.Mq, req, resp, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return err
		}

		return m.Rdb.FavoriteSet(req.UserId, req.VideoId, false)
	})
}

func favoriteDel(tx *mysql.DbTransaction, ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	_, err = tx.FavoriteDel(req.UserId, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
		return err
	} else if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	return nil
}

