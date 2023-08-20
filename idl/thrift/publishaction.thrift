// 发布视频请求
struct DouyinPublishActionRequest {
  1: required string Token; // 用户鉴权token
  2: required binary Data; // 视频数据
  3: required string Title; // 视频标题  
}

// 发布视频响应
struct DouyinPublishActionResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
}
