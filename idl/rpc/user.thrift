namespace go user
include "../base/rpc.thrift"
include "../base/http.thrift"

// 登录请求
struct douyin_user_login_request {
  1: string username
  2: string password
}

// 登录响应
struct douyin_user_login_response {
  1: i32 status_code
  2: string status_msg
  3: i64 user_id
  4: string token  
}

// 注册请求
struct douyin_user_register_request {
  1: string username
  2: string password
}

// 注册响应
struct douyin_user_register_response {
  1: i32 status_code
  2: string status_msg
  3: i64 user_id
  4: string token
}

// 获取用户信息请求
struct douyin_user_request {
  1: i64 user_id
}

// 获取用户信息响应
struct douyin_user_response {
  1: i32 status_code
  2: string status_msg
  3: http.User user
}

service UserService {
  douyin_user_login_response Login(1: douyin_user_login_request req)

  douyin_user_register_response Register(1: douyin_user_register_request req)

  douyin_user_response UserInfo(1: douyin_user_request req)

  // 获取用户信息
  rpc.User GetUserInfo(1: i64 userId)

  // 更新获赞数量
  void UpdateFavoritedCount(1: i64 userId, 2: i64 newFavoritedCount)

  // 更新关注数和粉丝数
  void UpdateFollowingAndFollowerCount(1: i64 userId, 2: i64 newFollowingCount, 3: i64 newFollowerCount)

}
