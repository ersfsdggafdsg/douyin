// Code generated by hertz generator. DO NOT EDIT.

package relation

import (
	relation "douyin/cmd/api/biz/handler/relation"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			_relation.POST("/action", append(_relationactionMw(), relation.RelationAction)...)
			{
				_follow := _relation.Group("/follow", _followMw()...)
				_follow.GET("/list", append(_followlistMw(), relation.FollowList)...)
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				_follower.GET("/list", append(_followerlistMw(), relation.FollowerList)...)
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				_friend.GET("/list", append(_friendlistMw(), relation.FriendList)...)
			}
		}
	}
}