package mq

import (
	"context"
	"douyin/shared/config"
	"douyin/shared/utils/mq"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageQueueManager struct {
	Comment *mq.Queue
}

type publishCommentUpdateRecord struct {
	VideoId int64
	AddCount int64
}

func NewManager() MessageQueueManager {
	return MessageQueueManager {
		Comment: mq.NewQueue("cmt2pub", "cmt2pub"),
	}
}

func (m *MessageQueueManager) UpdatePublishCommentCount(videoId, addCount int64) error {
	req := publishCommentUpdateRecord{
		VideoId: videoId,
		AddCount: addCount,
	}
	
	json, err := sonic.Marshal(&req)
	if err != nil {
		// 这个大抵是不会有异常的
		klog.Error(err)
		return err
	}

	return m.Comment.Publish(json)
}

func (m *MessageQueueManager) updatePublishCommentCountConsume() {
	m.Comment.Consume(func(b []byte) error {
		klog.Debugf("publish recv: %s", string(b))
		var req publishCommentUpdateRecord
		err := sonic.Unmarshal(b, &req)
		if err != nil {
			klog.Error("Convert failed:", err)
			return err
		}

		return config.Clients.Publish.UpdateCommentCount(
			context.Background(), req.VideoId, req.AddCount)
	})
}


func (m *MessageQueueManager) RunConsumers() {
	go m.updatePublishCommentCountConsume()
}

