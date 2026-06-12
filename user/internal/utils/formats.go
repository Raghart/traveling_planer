package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
)

func FormatImageUrls(country string, imageUrls []string) string {
	var msgBuilder strings.Builder
	fmt.Fprintf(&msgBuilder, "# %s's Image\n", country)

	for idx, url := range imageUrls {
		fmt.Fprintf(&msgBuilder, "📷 **Image %d**\n\n", idx+1)
		fmt.Fprintf(&msgBuilder, "![Click Here!](%s) \n\n", url)
	}
	return msgBuilder.String()
}

func FormatHolidays(country string, countryFestivities []routing.FestivityData) string {
	var msgBuilder strings.Builder

	fmt.Fprintf(&msgBuilder, "# %s's Holidays \n", country)

	for _, festivity := range countryFestivities {
		fmt.Fprintf(&msgBuilder, "* **%s** - %s\n", festivity.Date, festivity.Name)
	}
	return msgBuilder.String()
}

func FormatDescription(countryName, countryDescription string) string {
	var msgBuilder strings.Builder
	fmt.Fprintf(&msgBuilder, "# %s's Description \n", countryName)
	fmt.Fprint(&msgBuilder, countryDescription)
	return msgBuilder.String()
}

func FormatCurrency(currencyInfo routing.Currency) string {
	var msgBuilder strings.Builder
	fmt.Fprintf(&msgBuilder, "# Exchange rates from %s to %s\n\n",
		currencyInfo.From, currencyInfo.To)

	fmt.Fprintf(&msgBuilder, "* **1 %s** = %f %s\n", currencyInfo.FromCurrency,
		currencyInfo.ExchangeRate, currencyInfo.ToCurrency)

	fmt.Fprintf(&msgBuilder, "* **1 %s** = %f %s\n", currencyInfo.ToCurrency,
		currencyInfo.InverseRate, currencyInfo.FromCurrency)

	return msgBuilder.String()
}

func FormatWeather(dailyTemps []routing.DailyTemp) string {
	var msgBuilder strings.Builder
	fmt.Fprintf(&msgBuilder, "# %d-day Weather Forecast \n\n", len(dailyTemps))
	today := time.Now()
	for i, tempData := range dailyTemps {
		timePassed := 24 * i
		dayTime := time.Duration(timePassed) * time.Hour
		day := today.Add(dayTime)

		if i == 0 {
			fmt.Fprint(&msgBuilder, formatWeatherMessage("Today", day, tempData))
			continue
		}
		fmt.Fprint(&msgBuilder, formatWeatherMessage(day.Weekday().String()[:3], day, tempData))
	}
	return msgBuilder.String()
}

func formatWeatherMessage(day string, date time.Time, dailyTemp routing.DailyTemp) string {
	var msgBuilder strings.Builder

	fmt.Fprintf(&msgBuilder, "* **%s**(%s): %s  **%.2f°C** / **%.2f°C**",
		day, date.GoString(),
		getWeatherEmoji(dailyTemp.WeatherCode),
		dailyTemp.Max, dailyTemp.Min)

	if dailyTemp.AparentMax > dailyTemp.Max+3 {
		fmt.Fprintf(&msgBuilder,
			" - Prepare for a **sunny** day! 🔅 The sensation will be **%.2f°C**",
			dailyTemp.AparentMax)
	}

	if dailyTemp.AparentMin < dailyTemp.AparentMin-3 {
		fmt.Fprintf(&msgBuilder,
			" - **Wrap up** with a coat to survive the freeze! ❄️ The sensation will be **%.2f°C**",
			dailyTemp.AparentMin)
	}

	fmt.Fprint(&msgBuilder, "\n")
	return msgBuilder.String()
}

func getWeatherEmoji(weatherCode int) string {
	switch {
	case weatherCode == 0:
		return "☀"
	case weatherCode <= 3:
		return "🌤"
	case weatherCode >= 61 && weatherCode <= 65:
		return "☔"
	default:
		return "☁"
	}
}
