package delivery

import (
	"context"
	"fmt"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/rabbitmq/amqp091-go"
)

func (c *msgDelivery) consumeRabbitMQ(queueName string) {

	conn, ch := driver.ConnectMQ(c.Cfg)
	defer driver.CloseMQ(conn, ch)

	msgs, err := ch.Consume(
		queueName,                             // queue
		fmt.Sprintf("consumer_%s", queueName), // consumer
		false,                                 // auto-ack
		false,                                 // exclusive
		false,                                 // no-local
		false,                                 // no-wait
		nil,                                   // argsW
	)
	if err != nil {
		c.Logger.Panicf("Failed to consume a queue: %v", err)
		return
	}

	k := make(chan bool)

	go func() {
		ctx := context.Background()
		for d := range msgs {
			if d.RoutingKey == models.QueueInsertLog {
				c.createLog(ctx, d, string(d.Body))
			}
		}
	}()

	c.Logger.Info("RabbitMQ Waiting for messages. To exit press CTRL+C")
	<-k

}

func (c *msgDelivery) createLog(ctx context.Context, d amqp091.Delivery, message string) {
	err := c.LogSrv.CreateByMessage(ctx, message)
	if err != nil {
		c.Logger.Warnf("Consume queue %s, failed to process	a message: %v", d.RoutingKey, err)
		d.Nack(false, false)
	} else {
		c.Logger.Infof("Consume queue %s, successfully processed:  %s", d.RoutingKey, d.Body)
		d.Ack(false)
	}
}
