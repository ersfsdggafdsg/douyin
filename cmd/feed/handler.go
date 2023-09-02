package main

import (
	"context"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/feed"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse)
	resp.NextTime = time.Now().UnixMilli()
	// LatestTime是毫秒表示的（客户端测试的结果）
	if req.LatestTime <= 0 {
		req.LatestTime = time.Now().UnixMilli()
	}
	
	infos, err := config.Clients.Publish.QueryRecentVideoInfos(ctx, req.LatestTime, 3)
	if err != nil {
		klog.Error("Can't get video list", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	resp.VideoList = make([]*base.Video, len(infos))
	
	for i, v := range infos {
		user, err := config.Clients.User.GetUserInfo(ctx, v.AuthorId)
		if err != nil {
			errno.BuildBaseResp(errno.SuccessCode, resp)
			return resp, nil
		}

		resp.VideoList[i] = rpc2http.Video(v)

		isFavorited, err := config.Clients.Favorite.IsFavorited(ctx, req.UserId, v.Id)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}
		resp.VideoList[i].IsFavorite = isFavorited

		user.IsFollow, err = config.Clients.Relation.IsFollowing(ctx, req.UserId, user.Id)
		if err != nil {
			klog.Info("Relation query failed:", err)
		}

		resp.VideoList[i].Author = rpc2http.User(user)
		if resp.NextTime > v.CreateTime {
			resp.NextTime = v.CreateTime
		}
	}
	return resp, nil
}
