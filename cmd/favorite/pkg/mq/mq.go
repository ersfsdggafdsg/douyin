package mq

import (
	"context"
	"douyin/shared/config"
	"douyin/shared/utils/mq"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageQueueManager struct {
	Publish *mq.Queue
	User    *mq.Queue
}

type publishFavoriteUpdateRecord struct {
	VideoId int64
	AddCount int64
}

func NewManager() MessageQueueManager {
	return MessageQueueManager {
		Publish: mq.NewQueue("fav2pub", "fav2pub"),
		User   : mq.NewQueue("fav2usr", "fav2usr"),
	}
}

func (m *MessageQueueManager) UpdatePublishFavorite(videoId, addCount int64) error {
	req := publishFavoriteUpdateRecord{
		VideoId: videoId,
		AddCount: addCount,
	}

	klog.Debug(req)
	
	json, err := sonic.Marshal(&req)
	if err != nil {
		// 这个大抵是不会有异常的
		klog.Error(err)
		return err
	}

	klog.Debugf("send: %s", json)

	return m.Publish.Publish(json)
}

func (m *MessageQueueManager) updatePublishFavoriteConsume() {
	m.Publish.Consume(func(b []byte) error {
		klog.Debugf("publish recv: %s", string(b))
		var req publishFavoriteUpdateRecord
		err := sonic.Unmarshal(b, &req)
		if err != nil {
			klog.Error("Convert failed:", err)
			return err
		}

		return config.Clients.Publish.UpdateFavoriteCount(
			context.Background(), req.VideoId, req.AddCount)
	})
}

type userFavoriteUpdateRecord struct {
	AuthorId int64
	UserId int64
	AddCount int64
}

func (m *MessageQueueManager) UpdateUserFavorite(authorId, userId, addCount int64) error {
	req := userFavoriteUpdateRecord{
		AuthorId: authorId,
		UserId: userId,
		AddCount: addCount,
	}
	
	json, err := sonic.Marshal(req)
	if err != nil {
		// 这个大抵是不会有异常的
		klog.Error(err)
		return err
	}

	klog.Debugf("send: %s", json)

	return m.User.Publish(json)
}

func (m *MessageQueueManager) updateUserFavoriteConsume() {
	m.User.Consume(func(b []byte) error {
		klog.Debugf("user recv: %s", string(b))
		var req userFavoriteUpdateRecord
		err := sonic.Unmarshal(b, &req)
		if err != nil {
			klog.Error("Convert failed:", err)
			return err
		}
		return config.Clients.User.UpdateFavoriteCount(
			context.Background(), req.AuthorId, req.UserId,
			req.AddCount)
	})
}

func (m *MessageQueueManager) RunConsumers() {
	go m.updateUserFavoriteConsume()
	go m.updatePublishFavoriteConsume()
}
