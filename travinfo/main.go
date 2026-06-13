package main

import (
	"fmt"
	"log"

	"github.com/Raghart/traveling_planer/travinfo/internal/pubsub"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
)

func main() {
	conn, ch, queue, msgs, err := pubsub.ConnectBunny()
	utils.FailsOnError(err, "unable to connect with server")

	defer conn.Close()
	defer ch.Close()
	forever := make(chan struct{})

	func() {
		for d := range msgs {
			var err error
			switch d.Type {
			case "currency":
				err = pubsub.PublishLatestCurrency(d, ch, queue)
			case "temperature":
				err = pubsub.PublishCountryTemperature(d, ch, queue)
			case "holidays":
				err = pubsub.PublishCountryHolidays(d, ch, queue)
			case "description":
				err = pubsub.PublishCountryDescription(d, ch, queue)
			case "images":
				err = pubsub.PublishCountryImage(d, ch, queue)
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}

			if err != nil {
				log.Print(err)
				d.Nack(false, false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
