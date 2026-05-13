package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting traveling information microservice!")
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailsOnError(err, "unable to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailsOnError(err, "unable to create a channel")
	defer ch.Close()

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
	utils.FailsOnError(err, "unable to creathe the queue")

	err = ch.Qos(1, 0, false)
	utils.FailsOnError(err, "unable to setup QoS")

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailsOnError(err, "unable to create the channel")
	forever := make(chan struct{})

	func() {
		for d := range msgs {
			switch d.Type {
			case "currency":
				currencyData := &types.Currency{}
				json.Unmarshal(d.Body, currencyData)

				err = utils.PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency", queue,
					currencyData)
				utils.FailsOnError(err, "unable to publish the json")

				d.Ack(false)
			default:
				log.Println("Invalid request made!")
			}
		}
	}()

	fmt.Println("Starting travinfo, ready to recieve requests!")
	<-forever
}
