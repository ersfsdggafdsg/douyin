package service

import (
	"context"
	"douyin/cmd/favorite/pkg/manager"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
)

func FavoriteList(m *manager.Manager, ctx context.Context, req *favorite.DouyinFavoriteListRequest, resp *favorite.DouyinFavoriteListResponse) (err error) {
	// 没有使用redis，因为是根据一个id遍历所有结果

	infos, err := m.Db.FavoriteList(req.UserId)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

	resp.VideoList = make([]*base.Video, len(infos))
	for i, v := range infos {
		videoInfo, err := config.Clients.Publish.VideoInfo(ctx, v.VideoId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return err
		}

		resp.VideoList[i] = rpc2http.Video(videoInfo)

		userInfo, err := config.Clients.User.GetUserInfo(ctx, videoInfo.AuthorId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return err
		}

		resp.VideoList[i].Author = rpc2http.User(userInfo)
		resp.VideoList[i].Author.IsFollow = false

		if videoInfo.AuthorId == req.UserId {
			// 自己不需要查自己是否关注自己
			continue
		}
		isFollowing, err := config.Clients.Relation.IsFollowing(ctx, req.UserId, videoInfo.AuthorId)
		if err != nil {
			// 不是很大的问题
			continue
		}
		
		resp.VideoList[i].Author.IsFollow = isFollowing

	}

	return nil
}
