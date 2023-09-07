package service

import (
	"context"
	"douyin/cmd/comment/pkg/mq"
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/utils/errno"
)

func updateRemote(q mq.MessageQueueManager, ctx context.Context, req *comment.DouyinCommentActionRequest, resp *comment.DouyinCommentActionResponse, addCount int64) error {
	err := q.UpdatePublishCommentCount(req.VideoId, addCount)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}
	return err
}
