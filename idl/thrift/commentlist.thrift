// 评论操作请求
struct DouyinCommentActionRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 VideoId; // 视频id
  3: required i32 ActionType; // 1-发布评论,2-删除评论
  4: optional string CommentText; // 用户填写的评论内容,在action_type=1的时候使用
  5: optional i64 CommentId; // 要删除的评论id,在action_type=2的时候使用
}

// 评论操作响应
struct DouyinCommentActionResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: optional Comment Comment; // 评论成功返回评论内容,不需要重新拉取整个列表
}

// 评论内容
struct Comment {
  1: required i64 Id; // 视频评论id
  2: required User User; // 评论用户信息
  3: required string Content;  // 评论内容
  4: required string CreateDate; // 评论发布日期,格式 mm-dd  
}

// 评论列表请求
struct DouyinCommentListRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 VideoId; // 视频id
}

// 评论列表响应
struct DouyinCommentListResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: list<Comment> CommentList; // 评论列表
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
