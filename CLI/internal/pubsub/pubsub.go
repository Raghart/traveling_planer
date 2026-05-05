package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
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

	queue, err := ch.QueueDeclare(queueName, queueType == DurableType, queueType != DurableType,
		queueType != DurableType, false, nil)

	if err != nil {
		return &ampq.Channel{}, ampq.Queue{}, fmt.Errorf("error while creating the queue: %v", err)
	}
	if err = ch.QueueBind(queueName, key, exchange, false, nil); err != nil {
		return &ampq.Channel{}, ampq.Queue{}, fmt.Errorf("error while binding the queue: %v", err)
	}
	return ch, queue, nil
}

type SimpleQueueType int

const (
	DurableType = iota
	TransientType
)
