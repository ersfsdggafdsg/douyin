namespace go publish
include "../base/http.thrift"

// 获取发布视频列表请求
struct douyin_publish_list_request {
	1: i64 user_id
	2: string token
}

// 获取发布视频列表响应
struct douyin_publish_list_response {
	1: i32 status_code
	2: string status_msg
	3: list<http.Video> video_list
}

// 发布视频请求
struct douyin_publish_action_request {
	1: string token
	2: binary data
	3: string title
}

// 发布视频响应
struct douyin_publish_action_response {
	1: i32 status_code
	2: string status_msg
}

service PublishService {
	// 获取发布视频列表
	douyin_publish_list_response PublishList(1: douyin_publish_list_request request)
		(api.get = "/douyin/publish/list/")

	// 发布视频
	douyin_publish_action_response PublishAction(1: douyin_publish_action_request request)
		(api.post = "/douyin/publish/action/")
}
