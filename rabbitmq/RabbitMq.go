package rabbitmq

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/streadway/amqp"
	"ims/logs"
	"sync"
	"time"
)

//生产者所需要实现的接口
type Producers interface {
	Push() string
}

//消费者所需要实现的接口
type Receivers interface {
	Consume([]byte) bool
}

//RabbitMq 结构体
type RabbitMq struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	queueName    string //队列名称
	routerKey    string // 路由key
	exchangeName string //交换机名称
	exchangeType string //交换机类型
	producerList []Producers
	receiverList []Receivers
	mu           sync.RWMutex
}

//交换机对象
type QueueExchange struct {
	QueueName    string //交换机名称
	RouterKey    string //路由key
	ExchangeName string //交换机名称
	ExchangeType string //交换机类型
}

// 链接rabbitMQ
func (self *RabbitMq)mqConnect() {
	var err error
	mqConf,err := config.NewConfig("ini","./conf/datasource.conf")
	if err != nil {
		logs.NewLogs().Print("datasource.conf 获取失败 获取不到rabbitmq信息")
		return
	}
	host := mqConf.String("rabbitmq::host")
	port := mqConf.String("rabbitmq::port")
	user := mqConf.String("rabbitmq::user")
	pass := mqConf.String("rabbitmq::password")
	linkUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",user,pass,host,port)
	MqConn , err = amqp.Dial(linkUrl)

	if err != nil {
		logs.NewLogs().Print("链接rabbitmq 失败")
		return
	}
	fmt.Println("channel")
	fmt.Println(self)


	self.conn =MqConn   // 赋值给RabbitMQ对象
	if err != nil {
		fmt.Printf("MQ打开链接失败:%s \n", err)
	}
	self.channel,err  = MqConn.Channel()  // 赋值给RabbitMQ对象
	if err != nil {
		fmt.Printf("MQ打开管道失败:%s \n", err)
	}

}

//新建一个 rabbitmq
func NewRabbitMq(q *QueueExchange) *RabbitMq {
	rabbitMq := &RabbitMq{}
	rabbitMq.exchangeName = q.ExchangeName
	rabbitMq.exchangeType = q.ExchangeType
	rabbitMq.routerKey = q.RouterKey
	rabbitMq.queueName = q.QueueName
	return rabbitMq
}

/**
关闭
*/
func (self *RabbitMq) Close() {
	var err error
	err = self.channel.Close()
	if err != nil {
		logs.NewLogs().Print("关闭channel 失败")
	}
	err = self.conn.Close()
	if err != nil {
		logs.NewLogs().Print("关闭mq链接 失败")
	}
}

/**
追加生成者
*/
func (self *RabbitMq) AddProducer(p Producers) {
	self.producerList = append(self.producerList, p)
}

/**
追加消费者
*/
func (self *RabbitMq) AddReceiver(r Receivers) {
	self.receiverList = append(self.receiverList, r)
}

func (self *RabbitMq) Run() {
	for _,value := range self.producerList {
		go self.sendProducer(value)
	}

	for _ ,value := range self.receiverList {
		go self.sendReceiver(value)
	}
	time.Sleep(1)
}

/**
	执行生成
 */
func (self *RabbitMq) sendProducer(p Producers)  {
	if self.channel == nil {
		self.mqConnect()
	}

	err :=self.channel.ExchangeDeclare(
		self.exchangeName,
		self.exchangeType,
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil)

	if err != nil {
		fmt.Println("ExchangeDeclare err")
		logs.NewLogs().Print(fmt.Sprintf("创建交换机失败:%s", err))
		return
	}

	err =  self.channel.Publish(self.exchangeName, self.routerKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(p.Push()),
	})

	if err != nil {
		logs.NewLogs().Print(fmt.Sprintf("MQ任务发送失败:%s", err))
		return
	}
}


/**
	执行消费
 */
func (self *RabbitMq) sendReceiver(r Receivers)  {
	if self.channel == nil {
		self.mqConnect()
	}
	_, err := self.channel.QueueDeclare(
		self.queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)

	if err != nil {
		logs.NewLogs().Print(fmt.Sprintf("创建队列失败:%s", err))
		return
	}
	err = self.channel.QueueBind(
		self.queueName, // name of the queue
		"",        // bindingKey
		self.exchangeName,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		logs.NewLogs().Print(fmt.Sprintf("绑定队列到交换机中失败:%s", err))
		return
	}
	deliveries, err := self.channel.Consume(
		self.queueName,
		"", //consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	for d := range deliveries {
		err := r.Consume(d.Body)
		if err {
			d.Ack(true)
		}else {
			self.sendReceiver(r)
		}
	}
}
