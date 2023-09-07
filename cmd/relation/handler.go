package main

import (
	"context"
	"douyin/cmd/relation/pkg/manager"
	"douyin/cmd/relation/pkg/service"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{
	manager.Manager
}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)
	// 自己不能关注自己，且帐号得是正数吧
	if req.ToUserId == req.UserId || req.ToUserId <= 0 || req.UserId <= 0 {
		errno.BuildBaseResp(errno.InvalidOperationCode, resp)
		return
	}

	switch req.ActionType {
	case 1:// 关注
		err = service.Follow(&s.Manager, ctx, req, resp)
		return resp, nil
	case 2:// 取关
		err = service.Unfollow(&s.Manager, ctx, req, resp)
		return resp, nil
	default:
		klog.Error("Unsupported action type")
		errno.BuildBaseResp(errno.InvalidActionTypeErrCode, resp)
		return resp, nil
	}
}

// FollowList implements the RelationServiceImpl interface.
// userid作为粉丝，查看正则关注的人
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp = new(relation.DouyinRelationFollowListResponse)
	err = service.FollowList(&s.Manager, ctx, req, resp)
	return
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp = new(relation.DouyinRelationFollowerListResponse)
	err = service.FollowerList(&s.Manager, ctx, req, resp)
	return
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// WARN: 可能对于“好友”的定义不是粉丝
	resp = new(relation.DouyinRelationFriendListResponse)
	err = service.FriendList(&s.Manager, ctx, req, resp)
	return resp, nil
}

// IsFollowing implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollowing(ctx context.Context, userId int64, fanId int64) (resp bool, err error) {
	return service.IsFollowing(&s.Manager, ctx, userId, fanId)
}
