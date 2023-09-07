package service

import (
	"context"
	"douyin/cmd/comment/pkg/manager"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
)
func CommentList(m *manager.Manager, ctx context.Context, req *comment.DouyinCommentListRequest, resp *comment.DouyinCommentListResponse) (err error) {
	// 为什么这里不使用redis？
	// 根据videoid去查评论信息，怎么说也得遍历一遍数据库吧，
	// 而且也没有rpc调用要读取单个的评论信息。
	// 但是，用户信息和关注这种是可以考虑用redis的。
	result, err := m.Db.CommentList(req.VideoId)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
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
			return err
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
	return nil
}
