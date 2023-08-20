// 关系操作请求
struct DouyinRelationActionRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 ToUserId; // 对方用户id
  3: required i32 ActionType; // 1-关注,2-取消关注
}

// 关系操作响应
struct DouyinRelationActionResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
}
