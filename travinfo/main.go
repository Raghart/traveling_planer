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
			switch d.Type {
			case "currency":
				err := pubsub.PublishLatestCurrency(d, ch, queue)
				if err != nil {
					log.Print(err)
				}
			case "temperature":
				err := pubsub.PublishCountryTemperature(d, ch, queue)
				if err != nil {
					log.Print(err)
				}
			case "holidays":
				err := pubsub.PublishCountryHolidays(d, ch, queue)
				if err != nil {
					log.Print(err)
				}
			case "description":
				err := pubsub.PublishCountryDescription(d, ch, queue)
				if err != nil {
					log.Print(err)
				}
			case "images":
				err := pubsub.PublishCountryImage(d, ch, queue)
				if err != nil {
					log.Print(err)
				}
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
