package main

import (
	"context"
	"douyin/cmd/favorite/pkg/manager"
	"douyin/cmd/favorite/pkg/service"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct {
	manager.Manager
}

// TODO: 增加中间件，查看是否是登录用户
// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)
	service.FavoriteList(&s.Manager, ctx, req, resp)
	return
}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)
	// VideoId一定是正数
	if req.VideoId <= 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
	}

	// 对不存在的视频，不能点赞
	info, err := config.Clients.Publish.VideoInfo(ctx, req.VideoId)
	if err != nil {
		// 查询出错
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	} else if info == nil {
		// 没有该视频
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return resp, nil
	}

	// 该视频存在
	switch req.ActionType {
	case 1:  // 点赞
		service.FavoriteAdd(&s.Manager, ctx, req, resp)
	case 2:  // 取消点赞
		service.FavoriteDel(&s.Manager, ctx, req, resp)
	default: // 参数错误
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
	}

	return resp, nil
}

func (s *FavoriteServiceImpl) IsFavorited(ctx context.Context, userId, videoId int64) (resp bool, err error) {
	return service.IsFavorited(&s.Manager, ctx, userId, videoId)
}
