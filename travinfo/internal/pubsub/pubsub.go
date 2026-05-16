package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	_ "github.com/joho/godotenv/autoload"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *ampq.Channel, exchange, key, msgID, msgType string,
	queue ampq.Queue, val T) error {

	body, err := json.Marshal(val)
	utils.FailsOnError(err, "unable to marshal the json value")
	return ch.PublishWithContext(context.Background(), exchange, key, false, false, ampq.Publishing{
		ContentType:   "application/json",
		CorrelationId: msgID,
		ReplyTo:       queue.Name,
		Type:          msgType,
		Body:          body,
	})
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
	currencyJson := &types.CurrencyJSON{}

	apiKey := os.Getenv("MONEYKEY")
	fmt.Println(apiKey)

	res, err := http.Get("https://api.fxratesapi.com/latest")
	utils.FailsOnError(err, "unable to contact with the API")

	resBody, err := io.ReadAll(res.Body)
	utils.FailsOnError(err, "response doesn't have a body")

	err = json.Unmarshal(resBody, currencyJson)
	utils.FailsOnError(err, "unable to unmarshal json body")

	dict, err := utils.StructToDict(currencyJson.ConversionRates)
	utils.FailsOnError(err, "unable to pack the dict")

	searchValue := dict[string(currencyData.To)]

	if value, isFloat := searchValue.(float64); isFloat {
		currencyData.Value = value
	}

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency",
		queue, currencyData)
	utils.FailsOnError(err, "unable to publish currency to JSON")

	d.Ack(false)
}
