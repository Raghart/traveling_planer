package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	ampq "github.com/rabbitmq/amqp091-go"
)

func DeliverLatestCurrency(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) {
	currencyData := &types.Currency{}
	err := json.Unmarshal(d.Body, currencyData)

	FailsOnError(err, "unable to publish the json")
	currencyJson := &types.CurrencyJSON{}

	res, err := http.Get("https://api.fxratesapi.com/latest")
	FailsOnError(err, "unable to contact with the API")

	resBody, err := io.ReadAll(res.Body)
	FailsOnError(err, "response doesn't have a body")

	err = json.Unmarshal(resBody, currencyJson)
	FailsOnError(err, "unable to unmarshal json body")

	dict, err := StructToDict(currencyJson.Rates)
	FailsOnError(err, "unable to pack the dict")

	searchValue := dict[string(currencyData.To)]

	if value, isFloat := searchValue.(float64); isFloat {
		currencyData.Value = value
	}

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency",
		queue, currencyData)
	FailsOnError(err, "unable to publish currency to JSON")

	d.Ack(false)
}
