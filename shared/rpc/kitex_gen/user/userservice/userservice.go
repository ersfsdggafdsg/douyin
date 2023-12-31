// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	rpc "douyin/shared/rpc/kitex_gen/rpc"
	user "douyin/shared/rpc/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":               kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"Register":            kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"UserInfo":            kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"GetUserInfo":         kitex.NewMethodInfo(getUserInfoHandler, newUserServiceGetUserInfoArgs, newUserServiceGetUserInfoResult, false),
		"UpdateFavoriteCount": kitex.NewMethodInfo(updateFavoriteCountHandler, newUserServiceUpdateFavoriteCountArgs, newUserServiceUpdateFavoriteCountResult, false),
		"UpdateFollowCount":   kitex.NewMethodInfo(updateFollowCountHandler, newUserServiceUpdateFollowCountArgs, newUserServiceUpdateFollowCountResult, false),
		"UpdateWorkCount":     kitex.NewMethodInfo(updateWorkCountHandler, newUserServiceUpdateWorkCountArgs, newUserServiceUpdateWorkCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.UserId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func updateFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateFavoriteCountArgs)

	err := handler.(user.UserService).UpdateFavoriteCount(ctx, realArg.AuthorId, realArg.UserId, realArg.AddCount)
	if err != nil {
		return err
	}

	return nil
}
func newUserServiceUpdateFavoriteCountArgs() interface{} {
	return user.NewUserServiceUpdateFavoriteCountArgs()
}

func newUserServiceUpdateFavoriteCountResult() interface{} {
	return user.NewUserServiceUpdateFavoriteCountResult()
}

func updateFollowCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateFollowCountArgs)

	err := handler.(user.UserService).UpdateFollowCount(ctx, realArg.UserId, realArg.FanId, realArg.AddCount)
	if err != nil {
		return err
	}

	return nil
}
func newUserServiceUpdateFollowCountArgs() interface{} {
	return user.NewUserServiceUpdateFollowCountArgs()
}

func newUserServiceUpdateFollowCountResult() interface{} {
	return user.NewUserServiceUpdateFollowCountResult()
}

func updateWorkCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateWorkCountArgs)

	err := handler.(user.UserService).UpdateWorkCount(ctx, realArg.UserId, realArg.AddCount)
	if err != nil {
		return err
	}

	return nil
}
func newUserServiceUpdateWorkCountArgs() interface{} {
	return user.NewUserServiceUpdateWorkCountArgs()
}

func newUserServiceUpdateWorkCountResult() interface{} {
	return user.NewUserServiceUpdateWorkCountResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, userId int64) (r *rpc.UserInfo, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.UserId = userId
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateFavoriteCount(ctx context.Context, authorId int64, userId int64, addCount int64) (err error) {
	var _args user.UserServiceUpdateFavoriteCountArgs
	_args.AuthorId = authorId
	_args.UserId = userId
	_args.AddCount = addCount
	var _result user.UserServiceUpdateFavoriteCountResult
	if err = p.c.Call(ctx, "UpdateFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) UpdateFollowCount(ctx context.Context, userId int64, fanId int64, addCount int64) (err error) {
	var _args user.UserServiceUpdateFollowCountArgs
	_args.UserId = userId
	_args.FanId = fanId
	_args.AddCount = addCount
	var _result user.UserServiceUpdateFollowCountResult
	if err = p.c.Call(ctx, "UpdateFollowCount", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) UpdateWorkCount(ctx context.Context, userId int64, addCount int64) (err error) {
	var _args user.UserServiceUpdateWorkCountArgs
	_args.UserId = userId
	_args.AddCount = addCount
	var _result user.UserServiceUpdateWorkCountResult
	if err = p.c.Call(ctx, "UpdateWorkCount", &_args, &_result); err != nil {
		return
	}
	return nil
}
