// 点赞操作请求
struct DouyinFavoriteActionRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 VideoId; // 视频id
  3: required i32 ActionType; // 1-点赞,2-取消点赞
}

// 点赞操作响应
struct DouyinFavoriteActionResponse {
  1: required i32 StatusCode;  // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
}
