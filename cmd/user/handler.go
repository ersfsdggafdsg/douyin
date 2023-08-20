package main

import (
	rpc "/home/afeather/Codes/golang/src/douyin/shared/rpc/kitex_gen/rpc"
	user "/home/afeather/Codes/golang/src/douyin/shared/rpc/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, userId int64) (resp *rpc.UserInfo, err error) {
	// TODO: Your code here...
	return
}

// UpdateFavoritedCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFavoritedCount(ctx context.Context, userId int64, newFavoritedCount_ int64) (err error) {
	// TODO: Your code here...
	return
}

// UpdateFollowingAndFollowerCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFollowingAndFollowerCount(ctx context.Context, userId int64, newFollowingCount_ int64, newFollowerCount_ int64) (err error) {
	// TODO: Your code here...
	return
}
