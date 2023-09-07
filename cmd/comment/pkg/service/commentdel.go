package service

import (
	"context"
	"douyin/cmd/comment/pkg/manager"
	"douyin/cmd/comment/pkg/dal/mysql"
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils/errno"
	"errors"

	"gorm.io/gorm"
)

func CommentDel(m *manager.Manager, ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) error {
	// 评论id只有正数
	if req.CommentId <= 0 {
		errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
		return errno.RecodeNotFound
	}

	return m.Db.RunTransaction(func(tx *mysql.DbTransaction) error {
		err := commentDel(tx, ctx, req, resp)
		if err != nil {
			return nil
		}

		return updateRemote(m.Mq, ctx, req, resp, -1)
	})

}

func commentDel(tx *mysql.DbTransaction, ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse) (err error) {
	_, err = tx.CommentDel(req.CommentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
		return err
	} else if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	return nil
}

