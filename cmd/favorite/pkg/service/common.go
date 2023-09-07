package service

import (
	"context"
	"douyin/cmd/favorite/pkg/mq"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
	"errors"

	"gorm.io/gorm"
)

func updateRemote(ctx context.Context, q *mq.MessageQueueManager, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse, addCount int64) (err error) {
	/* 使用消息队列后带来的问题。
	 * 1. 如果去更新数量的时候，视频被删除了，那么更新一定会失败的，
	 *    但是此时作者可能还存在，所以如果作者自己把视频的点赞数量
	 *    数一遍，就会发现对不上了。
	 *    这可能是小问题。硬是要解决，就去更新用户的信息吧。
	 *
	 * 2. 也是视频被删除了，用户这边记录了点赞数量，记录了点赞的视频，
	 *    视频被删除了，就意味着不能够展示给用户看了，那么点赞数量和总
	 *    共显示的已经点赞的视频数量会不一致。
	 *    可能需要搞一个“视频已删除”
	 *
	 * 3. 作者这个用户已经删除，但是视频还在。
	 *    这种情况对作者自己没有影响，因为他看不到自己的任何信息了。
	 *
	 * 总之，大概需要做一个“视频已经删除”来显示。
	 */

	// 对不存在的视频，不能点赞，顺便把视频信息拉取过来，因为要用作者ID
	info, err := config.Clients.Publish.VideoInfo(ctx, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询出错
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	} else if err != nil {
		// 没有该视频
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return err
	}

	err = q.UpdatePublishFavorite(req.VideoId, addCount)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}

	// 更新视频作者的被点赞数和用户的点赞数
	err = q.UpdateUserFavorite(info.AuthorId, req.UserId, addCount)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	
	return nil
}
