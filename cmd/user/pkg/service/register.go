package service

import (
	"context"
	"douyin/cmd/user/pkg/manager"
	"douyin/shared/rpc/kitex_gen/user"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func Register(m *manager.Manager, ctx context.Context, req *user.DouyinUserRegisterRequest, resp *user.DouyinUserRegisterResponse) (err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
	}

	_, err = m.Db.QueryByName(req.Username)
	switch err {
	case gorm.ErrRecordNotFound:
		return register(m, ctx, req, resp)
	case nil:
		// 查询成功，表示用户已经存在
		errno.BuildBaseResp(errno.UserAlreadyExistErrCode, resp)
		klog.Error("User existed", err)
		return nil
	default:
		klog.Error("Query user failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil
	}
}

func register(m *manager.Manager, ctx context.Context, req *user.DouyinUserRegisterRequest, resp *user.DouyinUserRegisterResponse) (err error) {
	// 由于目前只有单个写入，并且没有其他需要同步的数据
	// 故而没有使用事务。
	info, err := m.Db.UserAdd(req.Username, req.Password)
	resp.UserId = info.Id
	if err != nil {
		klog.Error("Can't create user", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil
	}

	resp.Token, err = utils.GenerateToken(info.Id)
	if err != nil {
		klog.Error("Can't generate token", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil
	}

	return nil
}
