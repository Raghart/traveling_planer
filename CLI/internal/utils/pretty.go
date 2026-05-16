package utils

import (
	"fmt"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
)

func PresentWeather(dailyTemps []routing.DailyTemp) {
	today := time.Now()
	for i := range len(dailyTemps) {
		//Hoy (Vie): ☀️ 29°C / 24°C - ¡Día perfecto para la playa!
		timePassed := 24 * i
		dayTime := time.Duration(timePassed) * time.Hour
		day := today.Add(dayTime)

		if i == 0 {
			msgFormatted := formatWeatherMessage("Today", dailyTemps[i])
			fmt.Println(msgFormatted)
			continue
		}

		msgFormatted := formatWeatherMessage(day.Weekday().String()[:3], dailyTemps[i])
		fmt.Println(msgFormatted)
	}
}

func formatWeatherMessage(day string, dailyTemp routing.DailyTemp) string {
	msg := fmt.Sprintf("%s: %s  %.2f°C / %.2f°C",
		day,
		getWeatherEmoji(dailyTemp.WeatherCode),
		dailyTemp.Max,
		dailyTemp.Min,
	)

	if dailyTemp.AparentMax > dailyTemp.Max+3 {
		msg += fmt.Sprintf(
			" - Prepare for a sunny day! 🔅 The sensation will be %.2f",
			dailyTemp.AparentMax,
		)
	}

	if dailyTemp.AparentMin < dailyTemp.AparentMin-3 {
		msg += fmt.Sprintf(
			" - Wrap up with a Coat to survive the freeze! ❄️ The sensation will be %.2f",
			dailyTemp.AparentMin,
		)
	}

	return msg
}

func getWeatherEmoji(weatherCode int) string {
	switch {
	case weatherCode == 0:
		return "☀️"
	case weatherCode <= 3:
		return "⛅"
	case weatherCode >= 61 && weatherCode <= 65:
		return "🌧️"
	default:
		return "☁️"
	}
}
