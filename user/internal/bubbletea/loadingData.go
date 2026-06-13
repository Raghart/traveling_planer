package bubbletea

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/routing"
)

func (m *Model) LoadCountryData() chan tea.Msg {
	countryDataCh := make(chan tea.Msg, 5)
	go func() {
		dailyTemps, err := pubsub.GetTemperature(m.country.Destination)
		countryDataCh <- routing.CountryData{
			DataType: "temp",
			Data:     dailyTemps,
			Error:    err,
		}
	}()
	go func() {
		currencyData, err := pubsub.GetCurrency(m.country.From, m.country.Destination)
		countryDataCh <- routing.CountryData{
			DataType: "currency",
			Data:     currencyData,
			Error:    err,
		}
	}()
	go func() {
		holidays, err := pubsub.GetHolidays(m.country.Destination)
		countryDataCh <- routing.CountryData{
			DataType: "holidays",
			Data:     holidays,
			Error:    err,
		}
	}()
	go func() {
		description, err := pubsub.GetCountryDescription(m.country.Destination)
		countryDataCh <- routing.CountryData{
			DataType: "description",
			Data:     description,
			Error:    err,
		}
	}()
	go func() {
		urlImages, err := pubsub.GetUrlImages(m.country.Destination)
		countryDataCh <- routing.CountryData{
			DataType: "images",
			Data:     urlImages,
			Error:    err,
		}
	}()
	return countryDataCh
}

func (m *Model) CheckPercentage() tea.Cmd {
	if m.progress.Percent() >= 1.0 {
		m.loadingData = false
		cmd := m.list.SetItems(m.GenerateProjectActions())
		return cmd
	}

	return nil
}

func ListenCountryData(dataChannel chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		msg := <-dataChannel
		return msg
	}
}
