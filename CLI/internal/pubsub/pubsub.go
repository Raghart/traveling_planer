package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
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

func TestingRPC() (res string) {
	conn, err := ampq.Dial(routing.ConnectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	rabbitCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("error while creating the channel: %v", err)
	}

	q, err := rabbitCh.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	utils.FailOnError(err, "couldn't declare the user queue")

	msgs, err := rabbitCh.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "couldn't register a consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err = rabbitCh.PublishWithContext(
		ctx,
		"",
		"rpc_queue",
		false,
		false,
		ampq.Publishing{
			ContentType:   "text/plain",
			CorrelationId: "1",
			ReplyTo:       q.Name,
			Body:          []byte("Who"),
		},
	)

	for d := range msgs {
		if d.CorrelationId == "1" {
			res = string(d.Body)
			break
		}
	}
	return
}

type SimpleQueueType int

const (
	DurableType = iota
	TransientType
)
