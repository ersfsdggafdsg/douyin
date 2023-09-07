package mq

import (
	"sync"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)


// 初始化RabbitMq客户端，服务端需要安装
func initConnection() *amqp.Connection {
	connection, err := amqp.Dial(
		"amqp://guest:guest@localhost:5672/",
	)
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	return connection
}


type Queue struct {
	// 每一个队列对应一个Channel
	// 并且由于线程不安全，一个connection只能给一个Channel用
	mutex sync.Mutex
	connection *amqp.Connection
	channel *amqp.Channel
	queueName string
	exchangeName string
}

func NewQueue(queueName string, exchangeName string) (q *Queue) {
	var err error
	connection := initConnection()
	q = new(Queue)
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
		"fanout",    // 广播模式，将消息发送到所有绑定的队列上
					 // NOTICE: 在目前看来任何模式都行
		true,        // 这里表示Rabbitmq重启后，不自动删除交换机
		false,       // 是否自动删除交换机，这里是不删除
		false,       // 是否是内置交换机，这里表示不是内置的
		false,       // 这里表示等待成功
		nil)         // 没有更多的参数了

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

func (q *Queue)Publish(data []byte) error {
	// 这么做的原因是，rabbitmq不是线程安全的。
	// TODO: 使用链接池来优化
	q.mutex.Lock()
	defer q.mutex.Unlock()
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
	return err
}

func (q *Queue)GetConsumer() (<-chan amqp.Delivery, error) {
	klog.Debugf("consumer for %s created", q.queueName)
	return q.channel.Consume(
		q.queueName,// 队列名
		"",         // 消费者名称。这个无所谓，所以放空串
		false,      // 表示不自动向生产者回复ACK
					// 这里选择false的原因是，一些更新操作可能失败
		false,      // 表示不是唯一的消费者，这里和上面的设置是对应的
		false,      // 表示自己可以消费自己发送的消息
		false,      // 进行阻塞等待
		nil)        // 没有更多的参数了
}

// 传入参数：
// 能够解析、处理Publish发送的数据。
// 解析且处理成功，返回nil，否则返回非nil。
// nil的话，会回复rabbitmq当前消息已经处理成功
// 其他情况，回复处理失败，且放回到队列中
func (q *Queue)Consume(handler func([]byte) error) {
	msgs, err := q.GetConsumer()
	if err != nil {
		klog.Fatal("Can't receive massages", err)
	}

	// WARN: 感觉这里有风险
	for data := range msgs {
		if handler(data.Body) == nil {
			// 这里表示回复rabbitmq，当前消息处理成功
			data.Ack(false)
		} else {
			// data.Nack(multiple, requeue bool)
			// 回复rabbitmq，消息处理失败， 
			// 下面这个表示只确认当前收到的消息，并且放回到队列
			data.Nack(false, true)
		}
	}
}
