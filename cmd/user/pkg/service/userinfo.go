package service

import (
	"context"
	"douyin/cmd/user/pkg/manager"
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/rpc/kitex_gen/user"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetUserInfo(m *manager.Manager, ctx context.Context, userId int64) (resp *rpc.UserInfo, err error) {
	info, err := m.Db.QueryById(userId)
	if err != nil {
		klog.Error("Query user failed", err)
		return nil, err
	}
	resp = &info.UserInfo
	return
}

func UserInfo(m *manager.Manager, ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)
	info, err := GetUserInfo(m, ctx, req.UserId)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil, err
	}
	resp.User = rpc2http.User(info)
	resp.User.IsFollow = false
	return resp, nil
}
