package main

import (
	publish "douyin/shared/rpc/kitex_gen/publish"
	rpc "douyin/shared/rpc/kitex_gen/rpc"
	"context"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, request *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, request *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateCommentCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateCommentCount(ctx context.Context, videoId int64, newCommentCount_ int64) (err error) {
	// TODO: Your code here...
	return
}

// UpdateFavoriteCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateFavoriteCount(ctx context.Context, videoId int64, newFavoriteCount_ int64) (err error) {
	// TODO: Your code here...
	return
}

// QueryRecentVideoInfos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) QueryRecentVideoInfos(ctx context.Context, startTime int64, limit int64) (resp []*rpc.VideoInfo, err error) {
	// TODO: Your code here...
	return
}

// VideoInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) VideoInfo(ctx context.Context, videoId int64) (resp *rpc.VideoInfo, err error) {
	// TODO: Your code here...
	return
}
