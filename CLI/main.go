package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Welcome to Traveling Planner!")
	defer fmt.Println("Good traveling!")

	conn, err := amqp.Dial(routing.ConnectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	rabbitCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("error while creating the channel: %v", err)
	}

	if err := pubsub.PublishJSON(rabbitCh, routing.ExchangePerilDirect, routing.TestingKey,
		routing.CountryData{IsCountry: true}); err != nil {
		log.Fatalf("error while publishing the json: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CMD > ")
		scanner.Scan()
		if scanner.Text() == "exit" {
			break
		}
	}
}
