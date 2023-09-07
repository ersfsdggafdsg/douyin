package service

import (
	"context"
	"douyin/cmd/relation/pkg/manager"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
)

// FollowerList implements the RelationServiceImpl interface.
func FollowerList(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationFollowerListRequest, resp *relation.DouyinRelationFollowerListResponse) (err error) {
	userIds, err := m.Db.FansList(req.UserId)
	if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}
	resp.UserList = make([]*base.User, len(userIds))
	for i, uid := range userIds {
		userInfo, err := config.Clients.User.GetUserInfo(ctx, uid)
		if err != nil {
			klog.Error("Get user info failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return nil
		}
		resp.UserList[i] = rpc2http.User(userInfo)
		// 查询粉丝列表，用户不一定关注了粉丝，
		// 复用代码，使用缓存来优化查询。
		isFollowing, err := IsFollowing(m, ctx, req.UserId, userInfo.Id) 
		if err != nil {
			klog.Error("Get follow info failed")
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return nil
		} else {
			resp.UserList[i].IsFollow = isFollowing
		}
	}
	return
}
