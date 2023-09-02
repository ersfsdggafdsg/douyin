package main

import (
	"context"
	"douyin/cmd/user/pkg/mysql"
	rpc "douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/rpc/kitex_gen/user"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	Db mysql.UserManager
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
	}
	info, err := s.Db.QueryByName(req.Username)
	if err != nil {
		klog.Error("Query user failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	} else if info == nil {
		errno.BuildBaseResp(errno.UserNotExistErrCode, resp)
		return resp, nil
	} else if info.Password != req.Password {
		errno.BuildBaseResp(errno.WrongPasswordCode, resp)
		return resp, nil
	}

	resp.Token, err = utils.GenerateToken(info.Id)
	if err != nil {
		klog.Error("Can't generate token", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	resp.UserId = info.Id
	errno.BuildBaseResp(errno.SuccessCode, resp)
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
	}

	info, err := s.Db.QueryByName(req.Username)
	if err != nil {
		klog.Error("Query user failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	} else if info != nil {
		errno.BuildBaseResp(errno.UserAlreadyExistErrCode, resp)
		klog.Error("User existed", err)
		return resp, nil
	}

	info, err = s.Db.UserAdd(req.Username, req.Password)
	resp.UserId = info.Id
	if err != nil {
		klog.Error("Can't create user", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	resp.Token, err = utils.GenerateToken(info.Id)
	if err != nil {
		klog.Error("Can't generate token", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)
	info, err := s.GetUserInfo(ctx, req.UserId)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}
	resp.User = rpc2http.User(info)
	resp.User.IsFollow = false
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, userId int64) (resp *rpc.UserInfo, err error) {
	info, err := s.Db.QueryById(userId)
	if err != nil {
		klog.Error("Query user failed", err)
		return nil, err
	}
	return &info.UserInfo, nil
}

// UpdateFavoritedCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFavoritedCount(ctx context.Context, userId int64, addCount int64) (err error) {
	return s.Db.UpdateFavoritedCount(userId, addCount)
}

// UpdateFollowingCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFollowingCount(ctx context.Context, userId int64, addCount int64) (err error) {
	return s.Db.UpdateFollowingCount(userId, addCount)
}

// UpdateFollowerCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFollowerCount(ctx context.Context, userId int64, addCount int64) (err error) {
	return s.Db.UpdateFollowerCount(userId, addCount)
}

// UpdateWorkCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateWorkCount(ctx context.Context, userId int64, addCount int64) (err error) {
	return s.Db.UpdateWorkCount(userId, addCount)
}
