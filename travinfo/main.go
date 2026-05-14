package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
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

				utils.FailsOnError(err, "unable to publish the json")
				currencyJson := &types.CurrencyJSON{}

				res, err := http.Get("https://api.fxratesapi.com/latest")
				utils.FailsOnError(err, "unable to contact with the API")

				resBody, err := io.ReadAll(res.Body)
				utils.FailsOnError(err, "response doesn't have a body")

				err = json.Unmarshal(resBody, currencyJson)
				utils.FailsOnError(err, "unable to unmarshal json body")

				dict, err := utils.StructToDict(currencyJson.Rates)
				utils.FailsOnError(err, "unable to pack the dict")

				searchValue := dict[string(currencyData.To)]

				if value, isFloat := searchValue.(float64); isFloat {
					currencyData.Value = value
				}

				err = utils.PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency",
					queue, currencyData)
				utils.FailsOnError(err, "unable to publish currency to JSON")

				d.Ack(false)
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
