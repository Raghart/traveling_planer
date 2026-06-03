package main

import (
	"fmt"
	"log"

	"github.com/Raghart/traveling_planer/travinfo/internal/pubsub"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
)

func main() {
	conn, ch, queue, msgs, err := utils.ConnectBunny()
	utils.FailsOnError(err, "unable to connect with server")
	defer conn.Close()
	defer ch.Close()
	forever := make(chan struct{})

	func() {
		for d := range msgs {
			switch d.Type {
			case "currency":
				pubsub.DeliverLatestCurrency(d, ch, queue)
			case "temperature":
				pubsub.DeliverCountryTemperature(d, ch, queue)
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
