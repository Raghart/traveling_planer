package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/Raghart/traveling_planer/travinfo/internal/utils"
)

func main() {
	conn, ch, queue, msgs := utils.ConnectBunny()
	defer conn.Close()
	defer ch.Close()
	forever := make(chan struct{})

	func() {
		for d := range msgs {
			switch d.Type {
			case "currency":
				currencyData := &types.Currency{}
				err := json.Unmarshal(d.Body, currencyData)

				utils.FailsOnError(err, "unable to publish the json")
				currencyJson := &types.CurrencyJSON{}

				res, err := http.Get("https://api.fxratesapi.com/latest")
				utils.FailsOnError(err, "unable to contact with the API")

				resBody, err := io.ReadAll(res.Body)
				utils.FailsOnError(err, "response doesn't have a body")

				err = json.Unmarshal(resBody, currencyJson)
				utils.FailsOnError(err, "unable to unmarshal json body")

				dict, err := utils.StructToDict(currencyJson.Rates)
				utils.FailsOnError(err, "unable to pack the dict")

				searchValue := dict[string(currencyData.To)]

				if value, isFloat := searchValue.(float64); isFloat {
					currencyData.Value = value
				}

				err = utils.PublishJSON(ch, "", d.ReplyTo, d.CorrelationId, "currency",
					queue, currencyData)
				utils.FailsOnError(err, "unable to publish currency to JSON")

				d.Ack(false)
			default:
				log.Println("Invalid request made!")
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Starting traveling information microservice!")
	<-forever
}
