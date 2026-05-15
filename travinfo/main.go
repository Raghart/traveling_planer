package main

import (
	"fmt"
	"log"

	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
)

func main() {
	conn, ch, queue, msgs := utils.ConnectBunny()
	defer conn.Close()
	defer ch.Close()
	forever := make(chan struct{})

	func() {
		for d := range msgs {
			switch d.Type {
			case "currency":
				utils.DeliverLatestCurrency(d, ch, queue)
			case "temperature":
				utils.DeliverCountryTemperature(d, ch, queue)
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
