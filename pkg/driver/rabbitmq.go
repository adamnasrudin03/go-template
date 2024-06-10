package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Body      string
	QueueName string
}

var (
	cfg    = configs.GetInstance()
	logger = Logger(cfg)
)

func (r *RabbitMQ) Publish() {

	conn, ch := ConnectMQ(cfg)
	defer CloseMQ(conn, ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := r.Body
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err := ch.PublishWithContext(ctx,
		"",          // exchange
		r.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		message,
	)

	if err != nil {
		logger.Errorf("Failed to publish a message: %v", err)
		return
	}

	logger.Infof("Sent Queue %s body %s ", r.QueueName, body)
}

func (r *RabbitMQ) Consume() {

	conn, ch := ConnectMQ(cfg)
	defer CloseMQ(conn, ch)

	msgs, err := ch.Consume(
		r.QueueName,                             // queue
		fmt.Sprintf("consumer_%s", r.QueueName), // consumer
		true,                                    // auto-ack
		false,                                   // exclusive
		false,                                   // no-local
		false,                                   // no-wait
		nil,                                     // args
	)
	if err != nil {
		logger.Errorf("Failed to consume a queue: %v", err)
		return
	}

	k := make(chan bool)

	go func() {
		for d := range msgs {
			logger.Printf("Keys: %s", d.RoutingKey)
			logger.Printf("message: %s", d.Body)
			d.Nack(false, false)
		}
	}()

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("2.Received a message: %s", d.Body)
	// 		// d.Ack(false)
	// 	}
	// }()

	logger.Info(" RabbitMQ Waiting for messages. To exit press CTRL+C")
	<-k
}
