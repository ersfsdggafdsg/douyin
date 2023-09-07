package service

import (
	"context"
	"douyin/cmd/comment/pkg/dal/model"
	"douyin/cmd/comment/pkg/dal/mysql"
	"douyin/cmd/comment/pkg/manager"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
)

func CommentAdd(m *manager.Manager, ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) error {
	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		info, err := commentAdd(tx, ctx, req, resp)
		if err != nil {
			return err
		}

		err = updateRemote(m.Mq, ctx, req, resp, +1)
		if err != nil {
			return err
		}

		return buildCommentAddResp(ctx, info, resp)
	})
}

func commentAdd(tx *mysql.DbTransaction, ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) (info *model.Comment, err error) {
	info, err = tx.CommentAdd(req.UserId, req.VideoId, req.CommentText)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return nil, err
	}

	return info, nil
}

func buildCommentAddResp(ctx context.Context, info *model.Comment, resp *comment.DouyinCommentActionResponse) error {
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
