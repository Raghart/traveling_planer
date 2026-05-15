package utils

import (
	"encoding/json"
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

func StructToDict(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	err = json.Unmarshal(data, &res)
	return res, nil
}

func ConnectBunny() (*ampq.Connection, *ampq.Channel, ampq.Queue, <-chan ampq.Delivery) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	FailsOnError(err, "unable to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailsOnError(err, "unable to create a channel")

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
	FailsOnError(err, "unable to creathe the queue")

	err = ch.Qos(1, 0, false)
	FailsOnError(err, "unable to setup QoS")

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	FailsOnError(err, "unable to create the channel")
	return conn, ch, queue, msgs
}

func FailsOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
