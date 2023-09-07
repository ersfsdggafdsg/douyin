package main

import (
	"context"
	"douyin/cmd/user/pkg/manager"
	"douyin/cmd/user/pkg/service"
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/rpc/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	manager.Manager
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	service.Login(&s.Manager, ctx, req, resp)
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	service.Register(&s.Manager, ctx, req, resp)
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	return service.UserInfo(&s.Manager, ctx, req)
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, userId int64) (resp *rpc.UserInfo, err error) {
	return service.GetUserInfo(&s.Manager, ctx, userId) 
}

// UpdateFavoritedCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFavoriteCount(ctx context.Context, authorId, userId, addCount int64) (err error) {
	return service.UpdateFavoriteCount(&s.Manager, ctx, authorId, userId, addCount)
}

// UpdateFollowingCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateFollowCount(ctx context.Context, userId, fanId, addCount int64) (err error) {
	return service.UpdateFollowCount(&s.Manager, ctx, userId, fanId, addCount)
}

// UpdateWorkCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateWorkCount(ctx context.Context, userId int64, addCount int64) (err error) {
	return service.UpdateWorkCount(&s.Manager, ctx, userId, addCount)
}
