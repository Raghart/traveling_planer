package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	_ "github.com/joho/godotenv/autoload"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgID, msgType string,
	queue ampq.Queue, val T) error {

	body, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("unable to marshal the json value: %w", err)
	}
	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Type:          msgType,
		Body:          body,
	})
}

func ConnectBunny() (*ampq.Connection, *ampq.Channel, ampq.Queue, <-chan ampq.Delivery, error) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create a channel: %w", err)
	}

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
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create the queue: %w", err)
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to setup QoS: %w", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, ampq.Queue{}, nil, fmt.Errorf("unable to create the channel: %w", err)
	}

	return conn, ch, queue, msgs, nil
}

func DeliverCountryTemperature(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) {
	countryTemp := &types.CountryTemp{}
	err := json.Unmarshal(d.Body, countryTemp)
	utils.FailsOnError(err, "unable to unmarshal the body")

	mapLocations := utils.LoadLocations()
	countryLocation := mapLocations[countryTemp.Country]
	tempURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&daily=temperature_2m_max,temperature_2m_min,weather_code,precipitation_probability_max,apparent_temperature_max,apparent_temperature_min", countryLocation.Latitude, countryLocation.Longitude)

	res, err := http.Get(tempURL)
	utils.FailsOnError(err, fmt.Sprintf("unable to get the %s temperature", countryTemp.Country))

	bodyData, err := io.ReadAll(res.Body)
	utils.FailsOnError(err, "unable to read the body")

	countryTempJSON := &types.CountryTempJSON{}

	err = json.Unmarshal(bodyData, countryTempJSON)
	utils.FailsOnError(err, "unable to unmarshal the JSON API data")

	for i := 0; i < len(countryTempJSON.Daily.ApparentTemperatureMax); i++ {
		countryTemp.DailyTemperatures = append(countryTemp.DailyTemperatures, types.DailyTemp{
			WeatherCode: countryTempJSON.Daily.WeatherCode[i],
			Max:         countryTempJSON.Daily.Temperature2MMax[i],
			Min:         countryTempJSON.Daily.Temperature2MMin[i],
			RainProb:    countryTempJSON.Daily.PrecipitationProbabilityMax[i],
			AparentMax:  countryTempJSON.Daily.ApparentTemperatureMax[i],
			AparentMin:  countryTempJSON.Daily.ApparentTemperatureMin[i],
		})
	}

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "temperature", queue, countryTemp)
	utils.FailsOnError(err, "unable to publish the JSON")

	d.Ack(false)
}

func DeliverLatestCurrency(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) {
	currencyData := &types.Currency{}
	err := json.Unmarshal(d.Body, currencyData)
	utils.FailsOnError(err, "unable to publish the json")

	apiKey := os.Getenv("MONEYKEY")
	if strings.TrimSpace(apiKey) == "" {
		utils.FailsOnError(fmt.Errorf("key hasn't been loaded: %s", apiKey), "key not valid")
	}

	res, err := http.Get(
		fmt.Sprintf(
			"https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, currencyData.From))
	utils.FailsOnError(err, "unable to contact with the weather API")

	resBody, err := io.ReadAll(res.Body)
	utils.FailsOnError(err, "response doesn't have a body")

	currencyJson := &types.CurrencyJSON{}
	err = json.Unmarshal(resBody, currencyJson)
	utils.FailsOnError(err, "unable to unmarshal json body")

	dict, err := utils.StructToDict(currencyJson.ConversionRates)
	utils.FailsOnError(err, "unable to pack the dict")

	searchValue := dict[string(currencyData.To)]
	floatVal, isFloat := searchValue.(float64)

	if !isFloat {
		utils.FailsOnError(fmt.Errorf("value: '%v' wasn't found in the dict", searchValue),
			"searched value is not avaible in the dict")
	}

	currencyData.Value = floatVal

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency", queue, currencyData)
	utils.FailsOnError(err, "unable to publish currency to JSON")

	d.Ack(false)
}
