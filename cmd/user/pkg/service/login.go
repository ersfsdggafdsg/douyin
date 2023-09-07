package service

import (
	"context"
	"douyin/cmd/user/pkg/manager"
	"douyin/shared/rpc/kitex_gen/user"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func Login(m *manager.Manager, ctx context.Context, req *user.DouyinUserLoginRequest, resp *user.DouyinUserLoginResponse) (err error) {
	// 这里目前只有读取，如果要写入，请使用事务！
	if len(req.Username) == 0 || len(req.Password) == 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
	}

	info, err := m.Db.QueryByName(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.UserNotExistErrCode, resp)
		return nil
	} else if err != nil {
		klog.Error("Query user failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil
	} else if info.Password != req.Password {
		errno.BuildBaseResp(errno.WrongPasswordCode, resp)
		return nil
	}

	resp.Token, err = utils.GenerateToken(info.Id)
	if err != nil {
		klog.Error("Can't generate token", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil
	}

	resp.UserId = info.Id
	errno.BuildBaseResp(errno.SuccessCode, resp)
	return nil
}
