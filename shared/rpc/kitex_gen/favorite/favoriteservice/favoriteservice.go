// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "douyin/shared/rpc/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newFavoriteServiceFavoriteListArgs, newFavoriteServiceFavoriteListResult, false),
		"IsFavorited":    kitex.NewMethodInfo(isFavoritedHandler, newFavoriteServiceIsFavoritedArgs, newFavoriteServiceIsFavoritedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteListArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteListResult)
	success, err := handler.(favorite.FavoriteService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteListArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteListArgs()
}

func newFavoriteServiceFavoriteListResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteListResult()
}

func isFavoritedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceIsFavoritedArgs)
	realResult := result.(*favorite.FavoriteServiceIsFavoritedResult)
	success, err := handler.(favorite.FavoriteService).IsFavorited(ctx, realArg.UserId, realArg.VideoId)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newFavoriteServiceIsFavoritedArgs() interface{} {
	return favorite.NewFavoriteServiceIsFavoritedArgs()
}

func newFavoriteServiceIsFavoritedResult() interface{} {
	return favorite.NewFavoriteServiceIsFavoritedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (r *favorite.DouyinFavoriteActionResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteActionArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (r *favorite.DouyinFavoriteListResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteListArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavorited(ctx context.Context, userId int64, videoId int64) (r bool, err error) {
	var _args favorite.FavoriteServiceIsFavoritedArgs
	_args.UserId = userId
	_args.VideoId = videoId
	var _result favorite.FavoriteServiceIsFavoritedResult
	if err = p.c.Call(ctx, "IsFavorited", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
