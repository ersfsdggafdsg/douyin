package main

import (
	"context"
	"douyin/cmd/publish/pkg/model"
	"douyin/cmd/publish/pkg/mysql"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/publish"
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/utils"
	"douyin/shared/utils/cover"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"errors"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{
	Db mysql.PublishManager
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = new(publish.DouyinPublishListResponse)

	result, err := s.Db.QueryByAuthor(req.UserId)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	info, err := config.Clients.User.GetUserInfo(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.UserNotExistErrCode, resp)
	} else if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	info.IsFollow = false
	resp.VideoList = make([]*base.Video, len(result))

	for i, v := range result {
		// 因为model.VideoInfo内嵌了rpc.VideoInfo
		// 所以可以直接转换为rpc.VideoInfo，具体用法就是下面这样的
		resp.VideoList[i] = rpc2http.Video(&v.VideoInfo)
		resp.VideoList[i].Author = rpc2http.User(info)
	}
	return
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp = new(publish.DouyinPublishActionResponse)
	if len(req.Title) == 0 || len(req.Data) == 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return resp, nil
	}

	// 用户是否存在
	_, err = config.Clients.User.GetUserInfo(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.UserNotExistErrCode, resp)
		return resp, nil
	}

	hash := utils.SHA256(req.Data)
	videoHash := "vid" + hash
	coverHash := "cvr" + hash
	if utils.IsExists(videoHash) {
		// 由于使用的是哈希，所以可以直接判断是否存在
		klog.Info("Video existsted, create record only")
	} else {
		// 上传视频，写:=后，实际上创建了一个该块内的变量
		err = utils.Upload(req.Data, videoHash)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// 对视频截图，当作封面
		pic, err := cover.GetCoverFromBytes(req.Data)
		if err != nil {
			klog.Error(err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// 上传封面
		err = utils.Upload(pic, coverHash)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}
	}

	err = config.Clients.User.UpdateWorkCount(ctx, req.UserId, +1)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	info := model.VideoInfo {
		VideoInfo: rpc.VideoInfo {
			AuthorId: req.UserId,
			PlayHash: videoHash,
			CoverHash: coverHash,
			FavoriteCount: 0,
			CommentCount: 0,
			Title: req.Title,
		},
	}
	s.Db.Create(&info)
	return resp, nil
}

// UpdateCommentCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateCommentCount(ctx context.Context, videoId int64, addCount int64) (err error) {
	return s.Db.UpdateCommentCount(videoId, addCount)
}

// UpdateFavoriteCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) UpdateFavoriteCount(ctx context.Context, videoId int64, addCount int64) (err error) {
	return s.Db.UpdateFavoriteCount(videoId, addCount)
}

// QueryRecentVideoInfos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) QueryRecentVideoInfos(ctx context.Context, startTime int64, limit int64) (resp []*rpc.VideoInfo, err error) {
	// 实现Feed获取视频的功能
	result, err := s.Db.QueryRecentVideoInfos(startTime, limit)
	if err != nil {
		return nil, err
	}

	resp = make([]*rpc.VideoInfo, len(result))
	for i, v := range result {
		resp[i] = &v.VideoInfo
		resp[i].CreateTime = v.CreatedAt.UnixMilli()
	}
	return
}

// VideoInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) VideoInfo(ctx context.Context, videoId int64) (resp *rpc.VideoInfo, err error) {
	info, err := s.Db.QueryById(videoId)
	klog.Debug(info, err)
	if err != nil {
		return nil, err
	}
	return &info.VideoInfo, nil
}
