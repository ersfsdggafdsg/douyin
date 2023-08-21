package main

import (
	"context"
	"douyin/cmd/publish/pkg/cover"
	"douyin/cmd/publish/pkg/model"
	"douyin/cmd/publish/pkg/mysql"
	"douyin/shared/config"
	publish "douyin/shared/rpc/kitex_gen/publish"
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/tools"
	"douyin/shared/tools/rpc2http"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, request *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	result, err := mysql.GetUserWorks(request.UserId)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
		return resp, err
	}
	info, err := config.Clients.User.GetUserInfo(ctx, request.UserId)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
		return resp, err
	}
	info.IsFollow = false
	for i, v := range result {
		// 因为model.VideoInfo内嵌了rpc.VideoInfo
		// 所以可以直接转换为rpc.VideoInfo，具体用法就是下面这样的
		resp.VideoList[i] = rpc2http.Video(&v.VideoInfo)
		resp.VideoList[i].Author = rpc2http.User(info)
	}
	return
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, request *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	name := "vid" + tools.SHA256(request.Data)
	playUrl, err := tools.Upload(request.Data, name)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
		return resp, err
	}
	pic, err := cover.GetCoverFromUrl(playUrl)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
		return resp, err
	}
	name = "cvr" + tools.SHA256(pic)
	coverUrl, err := tools.Upload(pic, name)
	info := model.VideoInfo {
		VideoInfo: rpc.VideoInfo {
			AuthorId: request.UserId,
			PlayUrl: playUrl,
			CoverUrl: coverUrl,
			FavoriteCount: 0,
			CommentCount: 0,
			Title: request.Title,
		},
	}
	mysql.Create(&info)
	return
}

// UpdateCommentCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateCommentCount(ctx context.Context, videoId int64, newCommentCount_ int64) (err error) {
	return mysql.UpdateCommentCount(videoId, newCommentCount_)
}

// UpdateFavoriteCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateFavoriteCount(ctx context.Context, videoId int64, newFavoriteCount_ int64) (err error) {
	return mysql.UpdateFavoriteCount(videoId, newFavoriteCount_)
}

// QueryRecentVideoInfos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) QueryRecentVideoInfos(ctx context.Context, startTime int64, limit int64) (resp []*rpc.VideoInfo, err error) {
	result, err := mysql.QueryRecentVideoInfos(startTime, limit)
	if err != nil {
		return nil, err
	}
	resp = make([]*rpc.VideoInfo, len(result))
	for i, v := range result {
		resp[i] = &v.VideoInfo
	}
	return 
}

// VideoInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) VideoInfo(ctx context.Context, videoId int64) (resp *rpc.VideoInfo, err error) {
	// TODO: Your code here...
	return mysql.GetVideoInfo(videoId)
}
