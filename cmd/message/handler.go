package main

import (
	"douyin/cmd/message/pkg/mysql"
	"douyin/shared/rpc/kitex_gen/common"
	"douyin/shared/rpc/kitex_gen/message"
	"douyin/shared/utils/errno"

	"context"

	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{
	Db mysql.MessageManager
}

// LatestMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) LatestMessage(ctx context.Context, senderId int64, receiverId int64) (resp *common.Message, err error) {
	msg, err := s.Db.LatestMessage(senderId, receiverId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &msg.Message, nil
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	/* TODO: 使用premsgtime查最新消息，但是premsgtime是客户端时间而不是服务器时间。
	 * 在redis中创建uidtouid消息缓存，存的内容是最新消息时间，
	 * 如果最新消息时间是在premsgtime前后十秒内，且确实没有新写入的消息，
	 * 就返回空表，否则再去查询
	 * 也可以这么考虑：
	 * 如果写入了新的消息，在redis中删除uidtouid的记录，
	 * 下一次再来读取就标记服务器记录的时间，后续进行读取的话，
	 * 就进行上面的检查
	 */
	resp = new(message.DouyinMessageChatResponse)
	infos, err := s.Db.MessageList(req.UserId, req.ToUserId, req.PreMsgTime)
	if err != nil {
		klog.Error("Can't get message list", err)
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return resp, nil
	}
	resp.MessageList = make([]*common.Message, len(infos))
	for i, v := range infos {
		resp.MessageList[i] = &v.Message
	}
	return resp, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	// 目前只有发送消息的功能
	resp = new(message.DouyinMessageActionResponse)
	resp.StatusCode = int32(errno.SuccessCode)
	_, err = s.Db.MessageAdd(req.UserId, req.ToUserId, req.Content)
	if err != nil {
		klog.Error(err)
		errno.BuildBaseResp(errno.NotMotifiedCode, resp)
	}
	return 
}
