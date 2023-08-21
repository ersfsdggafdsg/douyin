package main

import (
	"context"
	"douyin/cmd/comment/pkg/mysql"
	comment "douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/tools"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, request *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp.StatusCode = consts.StatusOK
	switch request.ActionType {
	case 1:
		err = mysql.CommentAdd(request.UserId,
			request.VideoId, request.CommentText)
		if err != nil {
			tools.BuildBaseResp(err, consts.StatusNotModified, resp)
		}
		return
	case 2:
		err = mysql.CommentDel(request.CommentId)
		if err != nil {
			tools.BuildBaseResp(err, consts.StatusNotModified, resp)
		}
		return
	default:
		tools.BuildInvalidActionTypeResp(consts.StatusBadRequest, resp, err)
		return
	}
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, request *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
