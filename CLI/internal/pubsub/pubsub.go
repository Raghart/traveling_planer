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
