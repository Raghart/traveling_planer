package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgID, msgType string,
	queue ampq.Queue, val T) error {

	body, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("unable to marshal the json value: %w", err)
	}
	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Type:          msgType,
		Body:          body,
	})
}

func ConnectBunny() (*ampq.Connection, *ampq.Channel, ampq.Queue, <-chan ampq.Delivery, error) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create a channel: %w", err)
	}

	queue, err := ch.QueueDeclare(
		"travinfo-queue",
		true,
		false,
		false,
		false,
		ampq.Table{
			ampq.QueueTypeArg: ampq.QueueTypeQuorum,
		},
	)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create the queue: %w", err)
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to setup QoS: %w", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create the channel: %w", err)
	}

	return conn, ch, queue, msgs, nil
}
