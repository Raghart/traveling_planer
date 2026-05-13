package utils

import (
	"context"
	"encoding/json"
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgType, msgID string,
	queue ampq.Queue, val T) error {

	body, err := json.Marshal(val)
	FailsOnError(err, "unable to marshal the json value")
	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Type:          msgType,
		Body:          body,
	})
}

func FailsOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
