// 抖音视频流服务
namespace go feed
include "../base/http.thrift"

// 视频流请求
struct douyin_feed_request {
	1: i64 latest_time // 可选参数,限制返回视频的最新投稿时间戳,精确到秒,不填表示当前时间
	2: string token // 可选参数,登录用户设置
}

// 视频流响应
struct douyin_feed_response {
	1: i32 status_code // 状态码,0-成功,其他值-失败
	2: string status_msg // 返回状态描述
	3: list<http.Video> video_list // 视频列表
	4: i64 next_time // 本次返回的视频中,发布最早的时间,作为下次请求时的latest_time
}

service FeedService {
	// 获取视频流
	douyin_feed_response Feed(1: douyin_feed_request request)
		(api.get = "/douyin/feed/")
} 
