package main

import (
	"context"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/feed"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/tools"
	"douyin/shared/tools/rpc2http"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, request *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp = new(feed.DouyinFeedResponse)
	infos, err := config.Clients.Publish.QueryRecentVideoInfos(ctx, request.LatestTime, request.UserId)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
		return resp, err
	}
	resp.VideoList = make([]*base.Video, len(infos))
	for i, v := range infos {
		user, err := config.Clients.User.GetUserInfo(ctx, v.AuthorId)
		if err != nil {
			tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
			return resp, err
		}
		resp.VideoList[i] = rpc2http.Video(v)
		user.IsFollow, err = config.Clients.Relation.IsFollowing(ctx, request.UserId, user.Id)
		if err != nil {
			klog.Debug("Relation.IsFollowing failed:", err)
		}
		resp.VideoList[i].Author = rpc2http.User(user)
	}
	return
}
