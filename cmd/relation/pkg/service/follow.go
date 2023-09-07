package service

import (
	"context"
	"douyin/cmd/relation/pkg/dal/mysql"
	"douyin/cmd/relation/pkg/manager"
	"douyin/cmd/relation/pkg/dal/redis"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func Follow(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) (err error) {
	isFollowing, err := queryCache(&m.Rdb, ctx, req, resp)
	switch err {
	case redis.Nil:// 缓存未命中
		return followUpdate(m, ctx, req, resp)
	case nil:// 成功读取缓存
		if isFollowing == true {
			errno.BuildBaseResp(errno.AlreadyFollowedCode, resp)
			return errno.AlreadyFollowed
		} else {
			// 更新
			return followUpdate(m, ctx, req, resp)
		}
	default:// 出错
		return err
	}
}

func followUpdate(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) error {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		err := follow(tx, ctx, req, resp)
		if err != nil {
			return err
		}
		
		// 先发送消息再更新缓存，理由是消息队列最终操作了数据库，
		// 数据库更加重要，缓存大不了过期了再读取。
		err = m.Mq.UpdateUserFollowCount(req.ToUserId, req.UserId, +1)
		if err != nil {
			return err
		}

		err = m.Rdb.FollowingSet(req.ToUserId, req.UserId, true)
		if err != nil {
			return err
		}

		return nil
	})
}

func follow(tx *mysql.DbTransaction, ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) error {
	_, err := tx.FollowInfo(req.UserId, req.ToUserId)
	switch err {
	case gorm.ErrRecordNotFound:
		// 该关系不存在(RecordNotFound)，那么就需要创建了。
		_, err = tx.RelationAdd(req.UserId, req.ToUserId)
		if err != nil {
			klog.Error("Create relation failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
		}

		return err
	case nil:
		// 这种情况一定是已经关注了
		klog.Debug("Already followed")
		errno.BuildBaseResp(errno.AlreadyFollowedCode, resp)
		return errno.AlreadyFollowed

	default:
		klog.Error("Service error:", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

}

