package pubsub

import (
	"context"
	"encoding/json"
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key string, val T) error {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}

	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType: "application/json",
		Body:        jsonBytes,
	})
}

func DeclareAndBind(conn *ampq.Connection,
	exchange, queueName, key string,
	queueType SimpleQueueType) (*ampq.Channel, ampq.Queue, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("error while creating the channel: %v", err)
	}

	durable := false
	autoDelete := false
	exclusive := false

	if queueType == DurableType {
		durable = true
	} else {
		autoDelete = true
		exclusive = true
	}
	queue, err := ch.QueueDeclare(queueName, durable, autoDelete, exclusive, false, nil)
	if err != nil {
		log.Fatalf("there was an error while creating the queue: %v", err)
	}
	if err = ch.QueueBind(queueName, key, exchange, false, nil); err != nil {
		log.Fatal(err)
	}
	return ch, queue, nil
}

type SimpleQueueType int

const (
	DurableType = iota
	TransientType
)
