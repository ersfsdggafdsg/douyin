package main

import (
	"context"
	"douyin/cmd/relation/pkg/mysql"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
	"errors"

	"gorm.io/gorm"

	"github.com/cloudwego/kitex/pkg/klog"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{
	Db mysql.RelationManager
}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)
	// 自己不能关注自己
	if req.ToUserId == req.UserId {
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return
	}

	switch req.ActionType {
	case 1:// 关注
		// 与favorite类似，也得查看是否已经点赞了。
		err = s.follow(ctx, req, resp)
		if err != nil {
			return resp, nil
		}

		// TODO: 更改为消息队列
		err = config.Clients.User.UpdateFollowerCount(ctx, req.ToUserId, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// TODO: 更改为消息队列
		err = config.Clients.User.UpdateFollowingCount(ctx, req.UserId, +1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		return resp, nil
	case 2:// 取关
		err = s.unFollow(ctx, req, resp)
		if err != nil {
			return resp, nil
		}

		// TODO: 更改为消息队列
		err = config.Clients.User.UpdateFollowerCount(ctx, req.ToUserId, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		// TODO: 更改为消息队列
		err = config.Clients.User.UpdateFollowingCount(ctx, req.UserId, -1)
		if err != nil {
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		return resp, nil
	default:
		klog.Error("Unsupported action type")
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
		return resp, nil
	}
}

func (s *RelationServiceImpl) follow(ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) (err error) {
	resp = new(relation.DouyinRelationActionResponse)

	info, err := s.Db.FollowInfo(req.UserId, req.ToUserId)
	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && notFound == false {
		// 出错且不是因为没有找到信息
		klog.Error("Service error:", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	if info != nil {
		klog.Debug("Already followed")
		errno.BuildBaseResp(errno.AlreadyFollowedCode, resp)
		return err
	}

	// 该关系不存在(RecordNotFound)，那么就需要创建了。
	_, err = s.Db.RelationAdd(req.UserId, req.ToUserId)
	if err != nil {
		klog.Error("Create relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}
	return nil
}

func (s *RelationServiceImpl) unFollow(ctx context.Context, req *relation.DouyinRelationActionRequest, resp *relation.DouyinRelationActionResponse) (err error) {
	_, err = s.Db.RelationDel(req.UserId, req.ToUserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return err
	} else if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return err
	}
	return nil
}

// FollowList implements the RelationServiceImpl interface.
// userid作为粉丝，查看正则关注的人
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp = new(relation.DouyinRelationFollowListResponse)
	userIds, err := s.Db.FollowList(req.UserId)
	if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}

	resp.UserList = make([]*base.User, len(userIds))
	for i, uid := range userIds {
		info, err := config.Clients.User.GetUserInfo(ctx, uid)
		if err != nil {
			klog.Error("Get user info failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		resp.UserList[i] = rpc2http.User(info)
		// 既然是查询关注列表，那么肯定都关注了
		resp.UserList[i].IsFollow = true
	}
	return
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp = new(relation.DouyinRelationFollowerListResponse)
	userIds, err := s.Db.FansList(req.UserId)
	if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}
	resp.UserList = make([]*base.User, len(userIds))
	for i, uid := range userIds {
		info, err := config.Clients.User.GetUserInfo(ctx, uid)
		if err != nil {
			klog.Error("Get user info failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}
		resp.UserList[i] = rpc2http.User(info)
		// 查询粉丝列表，用户不一定关注了粉丝，好在是本地进行查询
		if info, err := s.Db.FollowInfo(req.UserId, info.Id); err != nil {
			klog.Error("Get follow info failed")
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		} else if info == nil {
			// 没找到信息，那就是没有关注
			resp.UserList[i].IsFollow = false
		} else {
			resp.UserList[i].IsFollow = true
		}
	}
	return
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// WARN: 可能对于“好友”的定义不是粉丝
	resp = new(relation.DouyinRelationFriendListResponse)
	userIds, err := s.Db.FollowList(req.UserId)
	if err != nil {
		klog.Error("Delete relation failed", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
	}
	resp.UserList = make([]*base.FriendUser, len(userIds))
	for i, uid := range userIds {
		info, err := config.Clients.User.GetUserInfo(ctx, uid)
		if err != nil {
			klog.Error("Get user info failed", err)
			errno.BuildBaseResp(errno.ServiceErrCode, resp)
			return resp, nil
		}

		user := rpc2http.User(info)
		msg, err := config.Clients.Message.LatestMessage(ctx, req.UserId, user.Id)
		if err != nil {
			klog.Error("Get latest message failed", err)
			continue
		}

		content := "无消息"
		if msg != nil {
			content = msg.Content
		}

		msgType := int64(1)
		// message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
		if msg != nil && msg.ToUserId == req.UserId {
			msgType = 0
		}
		// 这么做的原因是，IDL不支持继承，也没有golang那样的嵌入。
		resp.UserList[i] = &base.FriendUser {
			Id             : user.Id,
			Name           : user.Name,
			FollowCount    : user.FollowCount,
			FollowerCount  : user.FollowerCount,
			IsFollow       : user.IsFollow,
			Avatar         : user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature      : user.Signature,
			TotalFavorited : user.TotalFavorited,
			WorkCount      : user.WorkCount,
			FavoriteCount  : user.FavoriteCount,
			Message        : content,
			MsgType        : msgType,
		}

		// 既然是查询关注列表，那么肯定都关注了
		resp.UserList[i].IsFollow = true
	}
	return resp, nil
}

// IsFollowing implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollowing(ctx context.Context, userId int64, fanId int64) (resp bool, err error) {
	_, err = s.Db.FollowInfo(userId, fanId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
