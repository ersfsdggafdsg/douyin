package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	// 每一个队列对应一个Channel
	channel *amqp.Channel
	queueName string
	exchangeName string
}

func NewQueue(connection *amqp.Connection, queueName string, exchangeName string) (q Queue) {
	var err error
	q.channel, err = connection.Channel()
	if err != nil {
		klog.Fatal("Can't create channel", err)
	}
	_, err = q.channel.QueueDeclare(
		queueName,// 队列名
		true,     // 是否持久化。因为业务场景中，消息一般要保存到磁盘上，所以是true
		false,    // 最后一个消费者断开链接后，是否自动删除队列。
				  // 这里是不删除，所以是false
		false,    // (exclusive)该队列是否可以被多个链接以及该链接的channel使用。
				  // 设置为false表示可以
		false,    // (noWait)是否阻塞。设置为false表示阻塞
		nil)      // 表示没有更多的参数了
	if err != nil {
		klog.Fatal("Can't create queue", err)
	}

	err = q.channel.ExchangeDeclare(
		exchangeName,// 交换机名。这里就使用了队列名，
		"fanout",        // 广播模式，将消息发送到所有绑定的队列上
		true,            // 这里表示Rabbitmq重启后，不自动删除交换机
		false,           // 是否自动删除交换机，这里是不删除
		false,           // 是否是内置交换机，这里表示不是内置的
		false,           // 这里表示等待成功
		nil,             // 没有更多的参数了
	)

	if err != nil {
		klog.Fatal("Can't create switch", err)
	}

	err = q.channel.QueueBind(
		queueName,       // 队列名。表示该channel要使用该队列      
		"",              // routing key。fanout模式下，这个参数没有 
		exchangeName,// 交换机名
		false,           // 这里表示进行阻塞等待
		nil)             // 没有更多的参数了

	if err != nil {
		klog.Fatal("Can't bind queue", err)
	}

	q.queueName = queueName
	q.exchangeName = exchangeName
	return q
}

func (q *Queue)Close() {
	q.channel.Close()
}

func (q *Queue)Publish(data []byte) {
	err := q.channel.Publish(
		q.exchangeName,   // 交换机名
		"",               // fanout模式下，这个参数无效
		false,            // 表示如果找不到队列，会直接丢弃消息
		false,            // 表示如果无法路由，会直接丢弃消息
		amqp.Publishing { // 消息的实体
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body: data,
		})
	if err != nil {
		klog.Error("Can't publish msg", err)
	}
}

func (q *Queue)Consume(handler func([]byte)) {
	msgs, err := q.channel.Consume(
		q.queueName, // 队列名
		"",          // 消费者名称。这个无所谓，所以放空串
		true,        // 自动向生产者回复ACK
		false,       // 表示不是唯一的消费者，这里和上面的设置是对应的
		false,       // 表示自己可以消费自己发送的消息
		false,       // 进行阻塞等待
		nil,         // 没有更多的参数了
	)
	if err != nil {
		klog.Fatal("Can't receive massages", err)
	}
	for data := range msgs {
		handler(data.Body)
	}
}
