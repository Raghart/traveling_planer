package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := ampq.Dial(routing.ConnectionStr)
	utils.FailOnError(err, "error while trying to connect the client")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "error while creating the channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue",
		true,
		false,
		false,
		false,
		ampq.Table{
			ampq.QueueTypeArg: ampq.QueueTypeQuorum,
		},
	)
	utils.FailOnError(err, "error while trying to declare the queue")

	err = ch.Qos(1, 0, false)
	utils.FailOnError(err, "error creating the QoS")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "failed to register a consumer")
	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		for d := range msgs {
			err = ch.PublishWithContext(ctx,
				"",
				d.ReplyTo,
				false,
				false,
				ampq.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte("am I?"),
				},
			)
			utils.FailOnError(err, "service couldn't publish answer")

			d.Ack(false)
		}
	}()

	fmt.Println("Microservice active, awaiting users requests!")
	<-forever
}
