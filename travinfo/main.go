package main

import (
	"fmt"

	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting traveling information microservice!")
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailsOnError(err, "unable to connect to RabbitMQ")

	if conn != nil {
		fmt.Println("connection sucessful!")
	}
}
