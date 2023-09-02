namespace go relation
include "../base/http.thrift"

// 关注/取消关注请求
struct douyin_relation_action_request {
	1: string token
	2: i64 to_user_id
	3: i32 action_type // 1-关注,2-取消关注
}

// 关注/取消关注响应
struct douyin_relation_action_response {
	1: i32 status_code
	2: string status_msg 
}

// 获取关注列表请求
struct douyin_relation_follow_list_request {
	1: i64 user_id
	2: string token
}

// 获取关注列表响应
struct douyin_relation_follow_list_response {
	1: i32 status_code
	2: string status_msg
	3: list<http.User> user_list
}

// 获取粉丝列表请求
struct douyin_relation_follower_list_request {
	1: i64 user_id
	2: string token
} 

// 获取粉丝列表响应
struct douyin_relation_follower_list_response {
	1: i32 status_code
	2: string status_msg
	3: list<http.User> user_list
}

// 好友列表请求
struct douyin_relation_friend_list_request {
	1: i64 user_id
	2: string token
}

// 好友列表响应
struct douyin_relation_friend_list_response {
	1: i32 status_code
	2: string status_msg
	3: list<http.FriendUser> user_list  
}

service RelationService {
	// 关注或取关
	douyin_relation_action_response RelationAction(1: douyin_relation_action_request req)
		(api.post = "/douyin/relation/action/")

	// 关注列表
	douyin_relation_follow_list_response FollowList(1: douyin_relation_follow_list_request req)
		(api.get = "/douyin/relation/follow/list/")

	// 粉丝列表
	douyin_relation_follower_list_response FollowerList(1: douyin_relation_follower_list_request req)
		(api.get = "/douyin/relation/follower/list/")

	// 好友（粉丝）列表
	douyin_relation_friend_list_response FriendList(1: douyin_relation_friend_list_request req)
		(api.get = "/douyin/relation/friend/list/")
}
