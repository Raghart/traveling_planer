package pubsub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
	ampq "github.com/rabbitmq/amqp091-go"
)

func PublishCountryTemperature(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) error {
	countryTemp := &types.CountryTemp{}
	err := json.Unmarshal(d.Body, countryTemp)
	if err != nil {
		return fmt.Errorf("unable to unmarshal the body: %w", err)
	}

	mapLocations := utils.LoadLocations()
	countryLocation := mapLocations[countryTemp.Country]

	baseUrl := "https://api.open-meteo.com/v1/forecast"
	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%f", countryLocation.Latitude))
	params.Add("longitude", fmt.Sprintf("%f", countryLocation.Longitude))
	params.Add("daily", strings.Join([]string{
		"temperature_2m_max",
		"temperature_2m_min",
		"weather_code",
		"precipitation_probability_max",
		"apparent_temperature_max",
		"apparent_temperature_min",
	}, ","))

	res, err := http.Get(fmt.Sprintf("%s?%s", baseUrl, params.Encode()))
	if err != nil || res.StatusCode >= 400 {
		return fmt.Errorf("unable to contact the weather api temperature: %w", err)
	}

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("unable to read the body: %w", err)
	}

	countryTempJSON := &types.CountryTempJSON{}

	err = json.Unmarshal(bodyData, countryTempJSON)
	if err != nil {
		return fmt.Errorf("unable to unmarshal the JSON API data: %w", err)
	}

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
	if err != nil {
		return fmt.Errorf("unable to publish the json: %w", err)
	}

	return d.Ack(false)
}

func PublishLatestCurrency(d ampq.Delivery, ch *ampq.Channel, queue ampq.Queue) error {
	currencyData := &types.Currency{}
	err := json.Unmarshal(d.Body, currencyData)
	if err != nil {
		return fmt.Errorf("unable to publish the json: %w", err)
	}

	apiKey := os.Getenv("MONEYKEY")
	if strings.TrimSpace(apiKey) == "" {
		return fmt.Errorf("key hasn't been loaded: %s", apiKey)
	}

	currMap := utils.LoadCurrencyMap()
	fromCountryCurrency, isCountry := currMap[string(currencyData.From)]

	if !isCountry {
		return fmt.Errorf("Unknown country '%s'. It is not a valid america country",
			currencyData.From)
	}

	res, err := http.Get(
		fmt.Sprintf(
			"https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, fromCountryCurrency))
	if err != nil || res.StatusCode >= 400 {
		return fmt.Errorf("unable to contact with the currency api: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("response doesn't have a body: %w", err)
	}

	currencyJson := &types.CurrencyJSON{}
	err = json.Unmarshal(resBody, currencyJson)
	if err != nil {
		return fmt.Errorf("unable to unmarshal json body: %w", err)
	}

	dict, err := utils.StructToDict(currencyJson.ConversionRates)
	if err != nil {
		return fmt.Errorf("unable to pack the dict: %w", err)
	}

	toCountryCurrency, isCurr := currMap[string(currencyData.To)]
	if !isCurr {
		return fmt.Errorf("the country %s is not valid", currencyData.To)
	}

	searchedExchangeRate := dict[string(toCountryCurrency)]
	exchangeRate, isFloat := searchedExchangeRate.(float64)

	if !isFloat {
		return fmt.Errorf("value: '%v' wasn't found in the dict", searchedExchangeRate)
	}

	currencyData.FromCurrency = fromCountryCurrency
	currencyData.ToCurrency = toCountryCurrency
	currencyData.ExchangeRate = exchangeRate
	currencyData.InverseRate = 1 / exchangeRate

	err = PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency", queue, currencyData)
	if err != nil {
		return fmt.Errorf("unable to publish currency to json: %w", err)
	}

	return d.Ack(false)
}
