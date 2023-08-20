namespace go favorite
include "../base/http.thrift"

struct douyin_favorite_action_request {
1: string token; // 用户鉴权token
2: i64 video_id; // 视频id
3: i32 action_type; // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
1: i32 status_code; // 状态码，0-成功，其他值-失败
2: string status_msg; // 返回状态描述
}

struct douyin_favorite_list_request {
1: i64 user_id; // 用户id
2: string token; // 用户鉴权token
}

struct douyin_favorite_list_response {
1: i32 status_code; // 状态码，0-成功，其他值-失败
2: string status_msg; // 返回状态描述
3: list<http.Video> video_list; // 用户点赞视频列表
}

service FavoriteService {
	douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req)
		(api.post = "/douyin/favorite/action");
	douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req)
		(api.get = "/douyin/favorite/list");
}

