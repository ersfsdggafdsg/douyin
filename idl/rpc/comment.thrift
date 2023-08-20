namespace go comment
include "../base/http.thrift"
// 用户评论服务

// 发布和删除评论请求
struct douyin_comment_action_request {
  1: i64 user_id // 用户id
  2: i64 video_id // 视频id
  3: i32 action_type // 1-发布评论,2-删除评论
  4: string comment_text // 用户填写的评论内容,在action_type=1的时候使用
  5: i64 comment_id // 要删除的评论id,在action_type=2的时候使用
}

// 发布和删除评论响应
struct douyin_comment_action_response {
  1: i32 status_code // 状态码,0-成功,其他值-失败
  2: string status_msg // 返回状态描述
  3: http.Comment comment // 评论成功返回评论内容,不需要重新拉取整个列表
}

// 获取评论列表请求
struct douyin_comment_list_request {
  1: i64 user_id // 用户id
  2: i64 video_id // 视频id
}

// 获取评论列表响应
struct douyin_comment_list_response {
  1: i32 status_code // 状态码,0-成功,其他值-失败
  2: string status_msg // 返回状态描述 
  3: list<http.Comment> comment_list // 评论列表
}

// 评论服务
service CommentService {
  // 发布和删除评论
  douyin_comment_action_response CommentAction(1: douyin_comment_action_request request)
  
  // 获取评论列表
  douyin_comment_list_response CommentList(1: douyin_comment_list_request request) 

}

