// 发布列表请求
struct DouyinPublishListRequest {
  1: required i64 UserId; // 用户id
  2: required string Token; // 用户鉴权token
}

// 发布列表响应
struct DouyinPublishListResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: list<Video> VideoList; // 用户发布的视频列表
} 

// 视频信息
struct Video {
  1: required i64 Id; // 视频唯一标识
  2: required User Author; // 视频作者信息
  3: required string PlayUrl; // 视频播放地址
  4: required string CoverUrl; // 视频封面地址
  5: required i64 FavoriteCount; // 视频的点赞总数
  6: required i64 CommentCount; // 视频的评论总数
  7: required bool IsFavorite; // true-已点赞,false-未点赞
  8: required string Title; // 视频标题
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
