package service

import (
	"context"
	"douyin/cmd/favorite/pkg/manager"
	"douyin/cmd/favorite/pkg/dal/mysql"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
	"gorm.io/gorm"
)

func FavoriteAdd(m *manager.Manager, ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		err = favoriteAdd(tx, ctx, req, resp)
		if err != nil {
			return err
		}
		
		// 这个顺序的理由也是数据库的数据更加重要，缓存大不了等他过期。
		err = updateRemote(ctx, &m.Mq, req, resp, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return err
		}

		return m.Rdb.FavoriteSet(req.UserId, req.VideoId, true)
	})
}

func favoriteAdd(tx *mysql.DbTransaction, ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	_, err = tx.FavoriteInfo(req.UserId, req.VideoId)

	switch err {
	case gorm.ErrRecordNotFound:
		// 只有这种情况才需要添加信息
		_, err = tx.FavoriteAdd(req.UserId, req.VideoId)
		if err != nil {
			errno.BuildBaseResp(errno.NotMotifiedCode, resp)
		}
		return err
	case nil:
		errno.BuildBaseResp(errno.AlreadyLikedCode, resp)
		return errno.AlreadyFollowed
	default:
		return err
	}
}

