package service

import (
	"context"
	"douyin/cmd/relation/pkg/dal/mysql"
	"douyin/cmd/relation/pkg/dal/redis"
	"douyin/cmd/relation/pkg/manager"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func Unfollow(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) (err error) {
	isFollowing, err := queryCache(&m.Rdb, ctx, req, resp)
	switch err {
	case redis.Nil:// 缓存未命中
		return unfollowUpdate(m, ctx, req, resp)
	case nil:// 成功读取缓存
		if isFollowing == true {
			errno.BuildBaseResp(errno.AlreadyFollowedCode, resp)
			return errno.AlreadyFollowed
		} else {
			// 更新
			return unfollowUpdate(m, ctx, req, resp)
		}
	default:// 出错
		return err
	}

}

func unfollowUpdate(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) (err error) {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		err := unfollow(tx, ctx, req, resp)
		if err != nil {
			return err
		}
		
		// 先发送消息再更新缓存，理由是消息队列最终操作了数据库，
		// 数据库更加重要，缓存大不了过期了再读取。
		err = m.Mq.UpdateUserFollowCount(req.ToUserId, req.UserId, -1)
		if err != nil {
			return err
		}

		err = m.Rdb.FollowingSet(req.ToUserId, req.UserId, false)
		if err != nil {
			return err
		}

		return nil
	})
}

func unfollow(tx *mysql.DbTransaction, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) error {
	_, err := tx.RelationDel(req.UserId, req.ToUserId)
	switch err {
	case gorm.ErrRecordNotFound:
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return errno.InvalidOperation
	case nil:
		return nil
	default:
		klog.Error("Service error:", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

}
