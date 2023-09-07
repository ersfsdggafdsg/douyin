package mq

import (
	"context"
	"douyin/shared/config"
	"douyin/shared/utils/mq"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageQueueManager struct {
	Follow *mq.Queue
}

type publishFollowUpdateRecord struct {
	UserId int64
	FanId int64
	AddCount int64
}

func NewManager() MessageQueueManager {
	return MessageQueueManager {
		Follow: mq.NewQueue("rlt2usr", "rlt2usr"),
	}
}

func (m *MessageQueueManager) UpdateUserFollowCount(userId, fanId, addCount int64) error {
	req := publishFollowUpdateRecord{
		UserId  : userId,
		FanId   : fanId,
		AddCount: addCount,
	}
	
	json, err := sonic.Marshal(&req)
	if err != nil {
		// 这个大抵是不会有异常的
		klog.Error(err)
		return err
	}

	return m.Follow.Publish(json)
}

func (m *MessageQueueManager) updateUserFollowCountConsume() {
	m.Follow.Consume(func(b []byte) error {
		klog.Debugf("publish recv: %s", string(b))
		var req publishFollowUpdateRecord
		err := sonic.Unmarshal(b, &req)
		if err != nil {
			klog.Error("Convert failed:", err)
			return err
		}

		return config.Clients.User.UpdateFollowCount(
			context.Background(), req.UserId, req.FanId, req.AddCount)
	})
}


func (m *MessageQueueManager) RunConsumers() {
	go m.updateUserFollowCountConsume()
}


