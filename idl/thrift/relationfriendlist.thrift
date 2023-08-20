// 粉丝列表请求
struct DouyinRelationFollowerListRequest {
  1: required i64 UserId; // 用户id
  2: required string Token; // 用户鉴权token
}

// 粉丝列表响应  
struct DouyinRelationFollowerListResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: list<User> UserList; // 用户列表
}

// 用户信息
struct User {
  1: required i64 Id; // 用户id
  2: required string Name; // 用户名称
  3: optional i64 FollowCount; // 关注总数
  4: optional i64 FollowerCount; // 粉丝总数
  5: required bool IsFollow; // true-已关注,false-未关注
  6: optional string Avatar; // 用户头像
  7: optional string BackgroundImage; // 用户个人页顶部大图
  8: optional string Signature; // 个人简介
  9: optional i64 TotalFavorited; // 获赞数量
  10: optional i64 WorkCount; // 作品数量 
  11: optional i64 FavoriteCount; // 点赞数量
}

// 好友列表请求
struct DouyinRelationFriendListRequest {
  1: required i64 UserId; // 用户id
  2: required string Token; // 用户鉴权token
}

// 好友列表响应
struct DouyinRelationFriendListResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: list<FriendUser> UserList; // 用户列表
}

// 好友用户信息
struct FriendUser extends User {
  1: optional string Message; // 和该好友的最新聊天消息
  2: required i64 MsgType; // message消息的类型,0 => 当前请求用户接收的消息, 1 => 当前请求用户发送的消息
}
