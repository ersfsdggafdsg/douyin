package mq

import (
	"douyin/shared/initialize"
	"os"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type message struct {
	idx int
}


func TestMessageQueue(t *testing.T) {
	conn := initialize.InitMq()
	queue := NewQueue(conn, "test", "test")
	idx := 0
	n := 100000
	published := make([]int8, n)
	for idx < n {
		go func(i int) {
			json, err := sonic.Marshal(&message{idx})
			if err != nil {
				t.Log(err)
			}
			published[i] = 1
			queue.Publish([]byte(json))
		}(idx)
		idx++
	}
	wg := sync.WaitGroup{}
	wg.Add(n)
	go queue.Consume(
		func (data []byte) error {
			defer wg.Done()
			msg := message{}
			sonic.Unmarshal(data, &msg)
			if published[msg.idx] != 1 {
				t.Logf("发送错误，第%d位置设置为了%d", msg.idx, published[msg.idx])
			} else {
				published[msg.idx] = 2
			}
			return nil
		})
	wg.Wait()
	for i, v := range published {
		if v != 2 {
				t.Logf("接收错误，第%d位置设置为了%d", i, published[i])
		}
	}
}
