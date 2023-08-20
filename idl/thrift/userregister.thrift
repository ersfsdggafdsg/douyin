namespace go douyin.core

/**
* 用户注册请求
*/
struct douyin_user_register_request {
  1: required string username; // 注册用户名,最长32个字符
  2: required string password; // 密码,最长32个字符  
}

/**
* 用户注册响应
*/
struct douyin_user_register_response {
  1: required i32 status_code; // 状态码,0-成功,其他值-失败
  2: optional string status_msg; // 返回状态描述
  3: required i64 user_id; // 用户id
  4: required string token; // 用户鉴权token
}
