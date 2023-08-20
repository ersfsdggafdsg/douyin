package main

import (
	common "/home/afeather/Codes/golang/src/douyin/shared/rpc/kitex_gen/common"
	message "/home/afeather/Codes/golang/src/douyin/shared/rpc/kitex_gen/message"
	"context"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// LatestMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) LatestMessage(ctx context.Context, senderId int64, receiverId int64) (resp *common.Message, err error) {
	// TODO: Your code here...
	return
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, request *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, request *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
