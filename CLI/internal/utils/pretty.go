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

		fmt.Printf("%s: %s  %.2f°C / %.2f°C\n",
			day.Weekday().String()[:3],
			getWeatherEmoji(dailyTemps[i].WeatherCode),
			dailyTemps[i].Max,
			dailyTemps[i].Min,
		)
	}
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
