package pubsub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func DeliverCountryTemperature(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) {
	countryTemp := &types.CountryTemp{}
	err := json.Unmarshal(d.Body, countryTemp)
	utils.FailsOnError(err, "unable to unmarshal the body")

	mapLocations := utils.LoadLocations()
	countryLocation := mapLocations[countryTemp.Country]
	tempURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m", countryLocation.Latitude, countryLocation.Longitude)

	res, err := http.Get(tempURL)
	utils.FailsOnError(err, fmt.Sprintf("unable to get the %s temperature", countryTemp.Country))

	bodyData, err := io.ReadAll(res.Body)
	utils.FailsOnError(err, "unable to read the body")

	countryTempJSON := &types.CountryTempJSON{}

	err = json.Unmarshal(bodyData, countryTempJSON)
	utils.FailsOnError(err, "unable to unmarshal the JSON API data")

	countryTemp.Value = countryTempJSON.Current.Temperature2M

	err = utils.PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "temperature", queue, countryTemp)
	utils.FailsOnError(err, "unable to publish the JSON")

	d.Ack(false)
}
