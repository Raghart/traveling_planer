package pubsub

import (
	"encoding/json"
	"fmt"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
	"github.com/google/uuid"
)

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

func GetCurrency(fromCurr, toCurr string) (routing.Currency, error) {
	conn, ch, q, msgs, err := ConnectBunny()
	if err != nil {
		return routing.Currency{}, fmt.Errorf("unable to connect with RabbitMQ: %w", err)
	}

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()

	err = PublishJSON(ch, "", "travinfo-queue", "currency", corrID, q, routing.Currency{
		From: fromCurr,
		To:   toCurr,
	})
	if err != nil {
		return routing.Currency{}, fmt.Errorf("unable to ask for currency: %w", err)
	}

	var currencyStruct routing.Currency
	for d := range msgs {
		if d.CorrelationId == corrID {
			currencyData := &routing.Currency{}
			err = json.Unmarshal(d.Body, currencyData)
			if err != nil {
				return routing.Currency{},
					fmt.Errorf("unable to unmarshal the recieved data: %w", err)
			}
			currencyStruct = *currencyData
			break
		}
	}

	return currencyStruct, nil
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

func GetCountryDescription(country string) (description string) {
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
			description = countryData.Description
			break
		}
	}
	return
}

func GetUrlImages(country string) (urlImages []string) {
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
			urlImages = countryAsciiImg.ImageUrls
			break
		}
	}
	return
}
