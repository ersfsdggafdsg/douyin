package main

import (
	"context"
	"douyin/cmd/comment/pkg/manager"
	"douyin/cmd/comment/pkg/service"
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{
	manager.Manager
}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)
	if req.UserId <= 0 {
		// 未登录的用户怎么能够发评论？
		errno.BuildBaseResp(errno.AuthorizationFailedErrCode, resp)
		return resp, nil
	}

	/* 这一段是检测视频是否还存在。
	 * 不过后面觉得这一步多余，因为用户正看视频，
	 * 你告诉他视频已经删除了，这种更加影响体验。
	 * 而且查询一遍还挺花时间的。
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
	*/
	
	switch req.ActionType {
	case 1: // 发送评论
		/* 这里和favorite不同，因为一个用户可以给一个视频发多个评论，
		 * 但是却只能点一个赞。
		 * 不过最好的还是使用transaction
		 */
		service.CommentAdd(&s.Manager, ctx, req, resp)
		return resp, nil
	case 2: // 删除评论
		service.CommentDel(&s.Manager, ctx, req, resp)
		return resp, nil
	default:
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
		return resp, nil
	}
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)
	service.CommentList(&s.Manager, ctx, req, resp)
	return resp, nil
}
