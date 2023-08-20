// 聊天消息请求
struct DouyinMessageChatRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 ToUserId; // 对方用户id
  3: required i64 PreMsgTime; // 上次最新消息的时间(新增字段-apk更新中)
}

// 聊天消息响应
struct DouyinMessageChatResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
  3: list<Message> MessageList; // 消息列表  
}

// 单条消息
struct Message {
  1: required i64 Id; // 消息id
  2: required i64 ToUserId; // 该消息接收者的id
  3: required i64 FromUserId; // 该消息发送者的id 
  4: required string Content; // 消息内容
  5: optional string CreateTime; // 消息创建时间
}


// 发送消息请求
struct DouyinRelationActionRequest {
  1: required string Token; // 用户鉴权token
  2: required i64 ToUserId; // 对方用户id
  3: required i32 ActionType = 1; // 发送消息
  4: required string Content; // 消息内容
}

// 发送消息响应
struct DouyinRelationActionResponse {
  1: required i32 StatusCode; // 状态码,0-成功,其他值-失败
  2: optional string StatusMsg; // 返回状态描述
}
