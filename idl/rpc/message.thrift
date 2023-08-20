namespace go message
include "../base/http.thrift"
include "../base/common.thrift"

// 获取聊天消息列表请求
struct douyin_message_chat_request {
  1: i64 user_id // 用户id
  2: i64 to_user_id // 对方用户id
  3: i64 pre_msg_time // 上次最新消息的时间
}

// 获取聊天消息列表响应
struct douyin_message_chat_response {
  1: i32 status_code // 状态码,0-成功,其他值-失败
  2: string status_msg // 返回状态描述
  3: list<common.Message> message_list // 消息列表
} 

// 发送消息请求
struct douyin_message_action_request {
  1: i64 user_id // 用户id
  2: i64 to_user_id // 对方用户id
  3: i32 action_type // 操作类型,1-发送消息
  4: string content // 消息内容   
}

// 发送消息响应
struct douyin_message_action_response {
  1: i32 status_code // 状态码,0-成功,其他值-失败
  2: string status_msg // 返回状态描述
}

service MessageService {

  // 获取最新消息
  common.Message LatestMessage(1: i64 senderId, 2: i64 receiverId)

  // 获取聊天消息列表
  douyin_message_chat_response MessageList(1: douyin_message_chat_request request)

  // 发送消息
  douyin_message_action_response MessageAction(1: douyin_message_action_request request)

}
