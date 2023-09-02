package main

import (
	"context"
	"douyin/cmd/comment/pkg/mysql"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	comment "douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"errors"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{
	Db mysql.CommentManager
}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)
	if req.UserId <= 0 {
		// 未登录的用户怎么能够发评论？
		errno.BuildBaseResp(errno.AuthorizationFailedErrCode, resp)
		return resp, nil
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
	
	switch req.ActionType {
	case 1: // 发送评论
		/* 这里和favorite不同，因为一个用户可以给一个视频发多个评论，
		 * 但是却只能点一个赞。
		 * 不过最好的还是使用transaction
		 */
		// TODO: 更换为消息队列调用
		err = config.Clients.Publish.UpdateCommentCount(ctx, req.VideoId, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		err = s.commentAdd(ctx, req, resp)
		return resp, nil
	case 2: // 删除评论
		// 评论id只有正数
		if req.CommentId <= 0 {
			errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
			return
		}

		// TODO: 更换为消息队列调用
		err = config.Clients.Publish.UpdateCommentCount(ctx, req.VideoId, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		err = s.commentDel(ctx, req, resp)
		return resp, nil
	default:
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
		return resp, nil
	}
}

// 将Action分解为两个函数，方便扩展
func (s *CommentServiceImpl)commentDel(ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) (err error) {
	resp = new(comment.DouyinCommentActionResponse)
	_, err = s.Db.CommentDel(req.CommentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
		return err
	} else if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	return nil
}

func (s *CommentServiceImpl)commentAdd(ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) (err error) {
	info, err := s.Db.CommentAdd(req.UserId,
		req.VideoId, req.CommentText)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

	user, err := config.Clients.User.GetUserInfo(ctx, info.UserId)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

	resp.Comment = &base.Comment{
		Id: int64(info.ID),
		Content: info.Content,
		User: rpc2http.User(user),
		CreateDate: utils.Time2Str(info.CreatedAt),
	}

	return nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)
	result, err := s.Db.CommentList(req.VideoId)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	resp.CommentList = make([]*base.Comment, len(result))
	for i, c := range result {
		resp.CommentList[i] = &base.Comment {
			Id        : int64(c.ID),
			Content   : c.Content,
			CreateDate: utils.Time2Str(c.CreatedAt),
		}
		userInfo, err := config.Clients.User.GetUserInfo(
			ctx, c.UserId)
		if err != nil {
			errno.BuildBaseResp(errno.SuccessCode, resp)
			return resp, nil
		}

		resp.CommentList[i].User = rpc2http.User(userInfo)
		if req.UserId == c.UserId {
			// 不需要检查用户关注自己
			continue
		}

		isFollowing, err := config.Clients.Relation.IsFollowing(
			ctx, req.UserId, c.UserId)
		if err != nil {
			// 这个错误算是比较小的错误
			continue
		}

		resp.CommentList[i].User.IsFollow = isFollowing
	}
	return resp, nil
}
