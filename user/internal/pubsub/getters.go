package pubsub

import (
	"encoding/json"
	"fmt"

	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/google/uuid"
)

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

	currencyStruct := routing.Currency{}
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

func GetTemperature(country string) ([]routing.DailyTemp, error) {
	conn, ch, queue, msgs, err := ConnectBunny()
	if err != nil {
		return []routing.DailyTemp{}, fmt.Errorf("unable to connect to RabbitMQ: %w", err)
	}

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "temperature", corrID, queue, routing.CountryTemp{
		Country: country,
	})
	if err != nil {
		return []routing.DailyTemp{}, fmt.Errorf("unable to publish the temperature request: %w", err)
	}

	tempSlice := []routing.DailyTemp{}
	for d := range msgs {
		if d.CorrelationId == corrID {
			countryTemp := &routing.CountryTemp{}
			err = json.Unmarshal(d.Body, countryTemp)
			if err != nil {
				return []routing.DailyTemp{}, fmt.Errorf("unable to unmarshal the body: %w", err)
			}
			tempSlice = countryTemp.DailyTemperatures
			break
		}
	}

	return tempSlice, nil
}

func GetHolidays(country string) ([]routing.FestivityData, error) {
	conn, ch, queue, msgs, err := ConnectBunny()
	if err != nil {
		return []routing.FestivityData{}, fmt.Errorf("failed to connect with RabbitMQ: %w", err)
	}

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "holidays", corrID, queue, routing.CountryFestivities{
		Country: country,
	})
	if err != nil {
		return []routing.FestivityData{}, fmt.Errorf("unable to publish the json: %w", err)
	}

	festivities := []routing.FestivityData{}
	for d := range msgs {
		if d.CorrelationId == corrID {
			countryFestivities := &routing.CountryFestivities{}
			err := json.Unmarshal(d.Body, countryFestivities)
			if err != nil {
				return []routing.FestivityData{},
					fmt.Errorf("unable to unmarshal festivity body: %w", err)
			}
			festivities = countryFestivities.Festivities
			break
		}
	}

	return festivities, nil
}

func GetCountryDescription(country string) (string, error) {
	conn, ch, queue, msgs, err := ConnectBunny()
	if err != nil {
		return "", fmt.Errorf("unable to connect with RabbitMQ: %w", err)
	}

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "description", corrID, queue, routing.CountryDescription{
		Name: country,
	})
	if err != nil {
		return "", fmt.Errorf("unable to publish the data: %w", err)
	}

	var description string
	for d := range msgs {
		if d.CorrelationId == corrID {
			countryData := &routing.CountryDescription{}
			err := json.Unmarshal(d.Body, countryData)
			if err != nil {
				return "", fmt.Errorf("unable to unmarshal the data: %w", err)
			}
			description = countryData.Description
			break
		}
	}

	return description, nil
}

func GetUrlImages(country string) ([]string, error) {
	conn, ch, queue, msgs, err := ConnectBunny()
	if err != nil {
		return []string{}, fmt.Errorf("unable to connect to RabbitMQ: %w", err)
	}

	defer conn.Close()
	defer ch.Close()

	corrID := uuid.NewString()
	err = PublishJSON(ch, "", "travinfo-queue", "images", corrID, queue, routing.CountryAsciiImg{
		Name: country,
	})
	if err != nil {
		return []string{}, fmt.Errorf("unable to publish json: %w", err)
	}

	urlImages := []string{}
	for d := range msgs {
		if d.CorrelationId == corrID {
			countryAsciiImg := &routing.CountryAsciiImg{}
			err := json.Unmarshal(d.Body, countryAsciiImg)
			if err != nil {
				return []string{}, fmt.Errorf("unable to unmarshal the json: %w", err)
			}
			urlImages = countryAsciiImg.ImageUrls
			break
		}
	}

	return urlImages, nil
}
