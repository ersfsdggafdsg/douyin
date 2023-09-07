package service

import (
	"context"
	"douyin/cmd/relation/pkg/manager"
	"douyin/shared/config"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
	"douyin/shared/utils/rpc2http"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FriendList(m *manager.Manager, ctx context.Context, req *relation.DouyinRelationFriendListRequest, resp *relation.DouyinRelationFriendListResponse) (err error) {
	// WARN: 可能对于“好友”的定义不是粉丝
	userIds, err := m.Db.FansList(req.UserId)
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
			return nil
		}

		user := rpc2http.User(info)
		msg, err := config.Clients.Message.LatestMessage(ctx, req.UserId, user.Id)
		if err != nil {
			klog.Error("Get latest message failed", err)
			continue
		}

		// 默认给个无消息
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
	return nil
}

