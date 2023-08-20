package main

import (
	comment "/home/afeather/Codes/golang/src/douyin/shared/rpc/kitex_gen/comment"
	"context"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, request *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, request *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
