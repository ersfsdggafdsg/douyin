package main

import (
	"context"
	"douyin/cmd/comment/pkg/mysql"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	comment "douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/tools"
	"douyin/shared/tools/rpc2http"

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
	resp.StatusCode = consts.StatusOK
	result, err := mysql.CommentList(request.VideoId)
	if err != nil {
		tools.BuildBaseResp(err, consts.StatusNoContent, resp)
		return
	}
	resp.CommentList = make([]*base.Comment, len(result))
	for i, c := range result {
		resp.CommentList[i] = &base.Comment {
			Id        : c.VideoId,
			Content   : c.Content,
			CreateDate: tools.Time2Str(c.CreatedAt),
		}
		userInfo, err := config.Clients.User.GetUserInfo(
			ctx, request.UserId)
		if err != nil {
			tools.BuildBaseResp(err, consts.StatusInternalServerError, resp)
			return resp, err
		}
		resp.CommentList[i].User = rpc2http.User(userInfo)
		if request.UserId == c.UserId {
			// 不需要检查用户关注自己
			continue
		}
		isFollowing, err := config.Clients.Relation.IsFollowing(
			ctx, request.UserId, c.UserId)
		if err != nil {
			continue
		}
		resp.CommentList[i].User.IsFollow = isFollowing
	}
	return
}
