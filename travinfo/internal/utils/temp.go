package utils

import (
	"encoding/json"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	ampq "github.com/rabbitmq/amqp091-go"
)

func DeliverCountryTemperature(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) {
	countryTemp := &types.CountryTemp{}
	err := json.Unmarshal(d.Body, countryTemp)
	FailsOnError(err, "unable to unmarshal the body")

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "temperature", queue, countryTemp)
	FailsOnError(err, "unable to publish the JSON")
}
