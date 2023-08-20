namespace go common

// 发生错误时的回应
struct BaseResponse {
  1: i32 status_code
  2: string status_msg
}

// 聊天消息
struct Message {
  1: i64 id // 消息id
  2: i64 to_user_id // 消息接收用户id 
  3: i64 from_user_id // 消息发送用户id
  4: string content // 消息内容
  5: string create_time // 消息创建时间
}
