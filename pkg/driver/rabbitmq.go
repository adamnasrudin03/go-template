package driver

import (
	"context"
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Body      string
	QueueName string
}

var cfg = configs.GetInstance()

func (r *RabbitMQ) Publish() {

	conn, ch := ConnectMQ(cfg)
	defer CloseMQ(conn, ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := r.Body
	message := amqp.Publishing{
		ContentType: "text/plain",
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
		log.Printf("Failed to publish a message: %s", err.Error())
		return
	}
	log.Printf(" [x] Sent %s\n", body)
}

func (r *RabbitMQ) Consume() {

	conn, ch := ConnectMQ(cfg)
	defer CloseMQ(conn, ch)

	msgs, err := ch.Consume(
		r.QueueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Printf("Failed to consume a queue: %s", err.Error())
		return
	}

	k := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("1.Received a message: %s", d.Body)
			// d.Ack(false)
		}
	}()

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("2.Received a message: %s", d.Body)
	// 		// d.Ack(false)
	// 	}
	// }()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-k
}
