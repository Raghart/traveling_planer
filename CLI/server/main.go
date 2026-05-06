package main

import (
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
}
