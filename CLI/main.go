package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CMD > ")
		scanner.Scan()
		switch scanner.Text() {
		case "exit":
			return
		case "who":
			q, err := rabbitCh.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
			)
			utils.FailOnError(err, "couldn't declare the user queue")

			msgs, err := rabbitCh.Consume(
				q.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			utils.FailOnError(err, "couldn't register a consumer")

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			err = rabbitCh.PublishWithContext(
				ctx,
				"",
				"rpc_queue",
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: "1",
					ReplyTo:       q.Name,
					Body:          []byte("Who"),
				},
			)

			for d := range msgs {
				if d.CorrelationId == "1" {
					fmt.Println(string(d.Body))
				}
			}
		}
	}
}
