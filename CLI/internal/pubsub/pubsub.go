package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
	"github.com/google/uuid"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgID string, queue ampq.Queue, val T) error {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}

	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Body:          jsonBytes,
	})
}

func DeclareAndBind(conn *ampq.Connection,
	exchange, queueName, key string,
	queueType SimpleQueueType) (*ampq.Channel, ampq.Queue, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("error while creating the channel: %v", err)
	}

	queue, err := ch.QueueDeclare(queueName, queueType == DurableType, queueType != DurableType,
		queueType != DurableType, false, nil)

	if err != nil {
		return &ampq.Channel{}, ampq.Queue{}, fmt.Errorf("error while creating the queue: %v", err)
	}
	if err = ch.QueueBind(queueName, key, exchange, false, nil); err != nil {
		return &ampq.Channel{}, ampq.Queue{}, fmt.Errorf("error while binding the queue: %v", err)
	}
	return ch, queue, nil
}

func ConnectBunny() (*ampq.Connection, *ampq.Channel, ampq.Queue, <-chan ampq.Delivery) {
	conn, err := ampq.Dial(routing.ConnectionStr)
	if err != nil {
		log.Fatal(err)
	}

	rabbitCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("error while creating the channel: %v", err)
	}

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
	return conn, rabbitCh, q, msgs
}

func TestingRPC() (res string) {
	conn, rabbitCh, q, msgs := ConnectBunny()
	defer conn.Close()
	defer rabbitCh.Close()

	corrID := uuid.NewString()
	err := PublishJSON(rabbitCh, "", "rpc_queue", corrID, q, routing.CountryData{
		IsCountry: false,
	})
	utils.FailOnError(err, "error while trying to publish the client JSON")

	for d := range msgs {
		if d.CorrelationId == corrID {
			countryData := &routing.CountryData{}
			err = json.Unmarshal(d.Body, countryData)
			utils.FailOnError(err, "unable to unmarshal data")

			if countryData.IsCountry == true {
				res = "There is a country in the horizon!"
				break
			}
			res = "unknown"
			break
		}
	}
	return
}

func AskCurrency(fromCurr, toCurr string) (value float32) {
	conn, ch, q, msgs := ConnectBunny()
	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()

	err := PublishJSON(ch, "", "travinfo-queue", corrID, q, routing.Currency{
		From: fromCurr,
		To:   toCurr,
	})
	utils.FailOnError(err, "unable to ask for currency")
	func() {
		for d := range msgs {
			if d.CorrelationId == corrID {
				currencyData := &routing.Currency{}
				err = json.Unmarshal(d.Body, currencyData)
				utils.FailOnError(err, "unable to unmarshal the recieved data")
				value = currencyData.Value
				break
			}
		}
	}()
	return
}

type SimpleQueueType int

const (
	DurableType = iota
	TransientType
)
