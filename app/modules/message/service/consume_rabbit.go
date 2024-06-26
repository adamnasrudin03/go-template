package service

import (
	"context"
	"fmt"

	"github.com/adamnasrudin03/go-template/app/models"
	"github.com/adamnasrudin03/go-template/pkg/driver"
)

func (srv *MessageSrv) consumeRabbitMQ(queueName string) {
	const opName = "MessageService-consumeRabbitMQ"
	conn, ch := driver.ConnectMQ(srv.Cfg, opName)
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
		srv.Logger.Warnf("%s Failed to consume a queue: %v", opName, err)
		return
	}

	k := make(chan bool)

	go func() {
		ctx := context.Background()
		for d := range msgs {
			switch d.RoutingKey {
			case models.QueueInsertLog:
				srv.createLog(ctx, d, string(d.Body))
			default:
				srv.Logger.Warnf("%s Unknown queue: %s", opName, d.RoutingKey)
				d.Nack(false, false)
			}
		}
	}()

	srv.Logger.Infof("%s RabbitMQ Waiting for messages. To exit press CTRL+C", opName)
	<-k

}
