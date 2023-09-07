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

func FollowList(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationFollowListRequest, resp *relation.DouyinRelationFollowListResponse) (err error) {
	userIds, err := m.Db.FollowList(req.UserId)
	if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}

	resp.UserList = make([]*base.User, len(userIds))
	for i, uid := range userIds {
		info, err := config.Clients.User.GetUserInfo(ctx, uid)
		if err != nil {
			klog.Error("Get user info failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return nil
		}

		resp.UserList[i] = rpc2http.User(info)
		// 既然是查询关注列表，那么肯定都关注了
		resp.UserList[i].IsFollow = true
	}
	return
}

