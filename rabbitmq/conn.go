package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

var MqConn *amqp.Connection


func InitRabbitMq()  {
	queue := &QueueExchange{
		"test.rabbit",
		"rabbit",
		"demo",
		"direct",}
	test := &Test{}

	rabbitmq :=NewRabbitMq(queue)
	rabbitmq.AddProducer(test)
	rabbitmq.AddReceiver(test)
	rabbitmq.Run()
	fmt.Println(rabbitmq)
	
}