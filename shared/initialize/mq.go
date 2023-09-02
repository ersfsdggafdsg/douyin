package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)


// 初始化RabbitMq客户端，服务端需要安装
func InitMq() *amqp.Connection {
	connection, err := amqp.Dial(
		"amqp://guest:guest@localhost:5672/",
	)
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	return connection
}

