// 用户登录请求
struct DouyinUserLoginRequest {
  1: required string Username; // 登录用户名
  2: required string Password; // 登录密码
}

// 用户登录响应
struct DouyinUserLoginResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: required i64 UserId; // 用户id
  4: required string Token; // 用户鉴权token
}
