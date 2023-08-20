// 用户请求
struct DouyinUserRequest {
  1: required i64 UserId; // 用户id
  2: required string Token; // 用户鉴权token
}

// 用户响应
struct DouyinUserResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: required User User; // 用户信息
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
