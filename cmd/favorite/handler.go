package main

import (
	"context"
	"douyin/cmd/favorite/pkg/mysql"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/favorite"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
	"errors"
	"gorm.io/gorm"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{
	Db mysql.FavoriteManager
}

// TODO: 增加中间件，查看是否是登录用户
// FavoriteAction implements the FavoriteServiceImpl interface.
// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp = new(favorite.DouyinFavoriteListResponse)
	resp.StatusCode = int32(errno.SuccessCode)

	infos, err := s.Db.FavoriteList(req.UserId)
	if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}

	resp.VideoList = make([]*base.Video, len(infos))
	for i, v := range infos {
		videoInfo, err := config.Clients.Publish.VideoInfo(ctx, v.VideoId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		resp.VideoList[i] = rpc2http.Video(videoInfo)

		userInfo, err := config.Clients.User.GetUserInfo(ctx, videoInfo.AuthorId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		resp.VideoList[i].Author = rpc2http.User(userInfo)
		resp.VideoList[i].Author.IsFollow = false

		if videoInfo.AuthorId == req.UserId {
			// 自己不需要查自己是否关注自己
			continue
		}
		isFollowing, err := config.Clients.Relation.IsFollowing(ctx, req.UserId, videoInfo.AuthorId)
		if err != nil {
			// 不是很大的问题
		}
		
		resp.VideoList[i].Author.IsFollow = isFollowing

	}

	return
}

func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)
	// VideoId一定是正数
	if req.VideoId <= 0 {
		errno.BuildBaseResp(errno.ParamErrCode, resp)
		return
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

	// 该视频存在
	switch req.ActionType {
	case 1:  // 点赞
		/* 为什么先查数据库，再去更新？
		 *   最好的做法还是使用gorm的transaction，
		 *   任何一次远程调用都可以回滚操作，
		 *   但是，目前看来，如果不先检查是否已经点赞了，
		 *   而是先更新远程数据，后点赞，那么，
		 *   玩意实际情况是已经点赞了，远程更新后还得在撤销操作
		 */
		err := s.favoriteAdd(ctx, req, resp)
		if err != nil {
			return resp, nil
		}

		// TODO: 更换为消息队列调用
		err = config.Clients.Publish.UpdateFavoriteCount(ctx, req.VideoId, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// TODO: 更换为消息队列调用
		err = config.Clients.User.UpdateFavoritedCount(ctx, info.AuthorId, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		info, err = config.Clients.Publish.VideoInfo(ctx, req.VideoId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

	case 2:  // 取消点赞
		err = s.favoriteDel(ctx, req, resp)
		if err != nil {
			return resp, nil
		}

		// TODO: 更换为消息队列调用
		err = config.Clients.Publish.UpdateFavoriteCount(ctx, req.VideoId, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// TODO: 更换为消息队列调用
		err = config.Clients.User.UpdateFavoritedCount(ctx, info.AuthorId, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		info, err = config.Clients.Publish.VideoInfo(ctx, req.VideoId)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

	default: // 参数错误
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
	}

	return resp, nil
}

// 将Action分解为两个函数，出错了返回error
func (s *FavoriteServiceImpl) favoriteAdd(ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	info, err := s.Db.FavoriteInfo(req.UserId, req.VideoId)
	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && notFound == false {
		// 如果不是因为没有找到信息而返回错误的话
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	if info != nil {
		// 找到了一条点赞记录
		errno.BuildBaseResp(errno.AlreadyLikedCode, resp)
		return gorm.ErrDuplicatedKey
	}

	_, err = s.Db.FavoriteAdd(req.UserId, req.VideoId)
	if err != nil {
		errno.BuildBaseResp(errno.NotMotifiedCode, resp)
	}
	return
}

// 将Action分解为两个函数，出错了返回error
func (s *FavoriteServiceImpl) favoriteDel(ctx context.Context, req *favorite.DouyinFavoriteActionRequest, resp *favorite.DouyinFavoriteActionResponse) (err error) {
	_, err = s.Db.FavoriteDel(req.UserId, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.RecodeNotFoundCode, resp)
		return err
	} else if err != nil {
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	return nil
}

func (s *FavoriteServiceImpl) IsFavorited(ctx context.Context, userId, videoId int64) (resp bool, err error) {
	_, err = s.Db.FavoriteInfo(userId, videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
