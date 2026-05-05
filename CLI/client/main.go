package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/routing"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("starting testing client consumer...")
	conn, err := ampq.Dial(routing.ConnectionStr)
	if err != nil {
		log.Fatalf("error while trying to connect the client: %v", err)
	}

	username, err := ClientWelcome()
	if err != nil {
		log.Fatal(err)
	}
	_, queue, err := pubsub.DeclareAndBind(conn,
		routing.ExchangePerilDirect,
		fmt.Sprintf("%s.%s", routing.TestingKey, username),
		routing.TestingKey, pubsub.TransientType)

	if err != nil {
		log.Fatalf("there was a problem while trying to Bind the queue")
	}
	fmt.Printf("queue declared: %s!", queue.Name)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("RabbitMQ Connecion Closed")
}

func ClientWelcome() (string, error) {
	fmt.Println("Welcome to the Peril client!")
	fmt.Println("Please enter your username:")
	words := GetInput()
	if len(words) == 0 {
		return "", errors.New("you must enter a username. goodbye")
	}
	username := words[0]
	fmt.Printf("Welcome, %s!\n", username)
	PrintClientHelp()
	return username, nil
}

func PrintClientHelp() {
	fmt.Println("Possible commands:")
	fmt.Println("* move <location> <unitID> <unitID> <unitID>...")
	fmt.Println("    example:")
	fmt.Println("    move asia 1")
	fmt.Println("* spawn <location> <rank>")
	fmt.Println("    example:")
	fmt.Println("    spawn europe infantry")
	fmt.Println("* status")
	fmt.Println("* spam <n>")
	fmt.Println("    example:")
	fmt.Println("    spam 5")
	fmt.Println("* quit")
	fmt.Println("* help")
}

func GetInput() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}
