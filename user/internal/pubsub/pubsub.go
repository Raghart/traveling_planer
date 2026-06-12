package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgType, msgID string, queue ampq.Queue, val T) error {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("error while trying to marshal the value: %w", err)
	}

	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Body:          jsonBytes,
		Type:          msgType,
	})
}

func DeclareAndBind(conn *ampq.Connection,
	exchange, queueName, key string,
	queueType SimpleQueueType) (*ampq.Channel, ampq.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, ampq.Queue{}, fmt.Errorf("error while creating the channel: %w", err)
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

func ConnectBunny() (*ampq.Connection, *ampq.Channel, ampq.Queue, <-chan ampq.Delivery, error) {
	conn, err := ampq.Dial(os.Getenv("CONNECTRABBIT"))
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to connect with rabbit server: %w", err)
	}

	rabbitCh, err := conn.Channel()
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("error while creating the channel: %w", err)
	}

	q, err := rabbitCh.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("couldn't declare the user queue: %w", err)
	}

	msgs, err := rabbitCh.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("couldn't register a consumer: %w", err)
	}
	return conn, rabbitCh, q, msgs, nil
}

func GetTestingRPC() (res string) {
	conn, rabbitCh, q, msgs, err := ConnectBunny()
	utils.FailOnError(err, "unable to connect to bunny")

	defer conn.Close()
	defer rabbitCh.Close()

	corrID := uuid.NewString()
	err = PublishJSON(rabbitCh, "", "rpc_queue", "", corrID, q, routing.TestingData{
		IsCountry: false,
	})
	utils.FailOnError(err, "error while trying to publish the client json")

	for d := range msgs {
		if d.CorrelationId == corrID {
			countryData := &routing.TestingData{}
			err = json.Unmarshal(d.Body, countryData)
			utils.FailOnError(err, "unable to unmarshal data")

			res = "unknown"
			if countryData.IsCountry == true {
				res = "There is a country in the horizon!"
			}
			break
		}
	}
	return
}

func GetCurrency(fromCurr, toCurr string) (currencyStruct routing.Currency) {
	conn, ch, q, msgs, err := ConnectBunny()
	utils.FailOnError(err, "unable to load to bunny")

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()

	err = PublishJSON(ch, "", "travinfo-queue", "currency", corrID, q, routing.Currency{
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
				currencyStruct = *currencyData
				break
			}
		}
	}()
	return
}

func GetTemperature(country string) (tempSlice []routing.DailyTemp) {
	conn, ch, queue, msgs, err := ConnectBunny()
	utils.FailOnError(err, "unable to connect to bunny server")

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "temperature", corrID, queue, routing.CountryTemp{
		Country: country,
	})
	utils.FailOnError(err, "unable to publish the temperature request")

	func() {
		for d := range msgs {
			if d.CorrelationId == corrID {
				countryTemp := &routing.CountryTemp{}
				err = json.Unmarshal(d.Body, countryTemp)
				utils.FailOnError(err, "unable to unmarshal the body")
				tempSlice = countryTemp.DailyTemperatures
				break
			}
		}
	}()
	return
}

func GetHolidays(country string) (festivities []routing.FestivityData) {
	conn, ch, queue, msgs, err := ConnectBunny()
	utils.FailOnError(err, "failed to connect with rabbitMQ")

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "holidays", corrID, queue, routing.CountryFestivities{
		Country: country,
	})
	utils.FailOnError(err, "unable to publish the json")

	func() {
		for d := range msgs {
			if d.CorrelationId == corrID {
				countryFestivities := &routing.CountryFestivities{}
				err := json.Unmarshal(d.Body, countryFestivities)
				utils.FailOnError(err, "unable to unmarshal festivity body")
				festivities = countryFestivities.Festivities
				break
			}
		}
	}()
	return
}

func GetCountryDescription(country string) (description routing.CountryDescription) {
	conn, ch, queue, msgs, err := ConnectBunny()
	utils.FailOnError(err, "unable to connect with RabbitMQ")

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "description", corrID, queue, routing.CountryDescription{
		Name: country,
	})
	utils.FailOnError(err, "unable to publish the data")

	for d := range msgs {
		if d.CorrelationId == corrID {
			countryData := &routing.CountryDescription{}
			err := json.Unmarshal(d.Body, countryData)
			utils.FailOnError(err, "unable to unmarshal the data")
			description = *countryData
			break
		}
	}
	return
}

func GetUrlImages(country string) (asciiImg routing.CountryAsciiImg) {
	conn, ch, queue, msgs, err := ConnectBunny()
	utils.FailOnError(err, "unable to connect to RabbitMQ")

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "images", corrID, queue, routing.CountryAsciiImg{
		Name: country,
	})
	utils.FailOnError(err, "unable to publish json")

	for d := range msgs {
		if d.CorrelationId == corrID {
			countryAsciiImg := &routing.CountryAsciiImg{}
			err := json.Unmarshal(d.Body, countryAsciiImg)
			utils.FailOnError(err, "unable to unmarshal the json")
			asciiImg = *countryAsciiImg
			break
		}
	}
	return
}

type SimpleQueueType int

const (
	DurableType = iota
	TransientType
)
