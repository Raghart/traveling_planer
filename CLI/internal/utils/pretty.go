package utils

import (
	"fmt"
	"time"

	"github.com/Raghart/traveling_planer/internal/routing"
)

func PresentWeather(dailyTemps []routing.DailyTemp) {
	today := time.Now()
	for i := 1; i < len(dailyTemps)+1; i++ {
		//Hoy (Vie): ☀️ 29°C / 24°C - ¡Día perfecto para la playa!
		timePassed := 24 * i
		dayTime := time.Duration(timePassed) * time.Hour
		day := today.Add(dayTime)

		fmt.Printf("%s: %s %f°C / %f°C\n",
			day.Weekday().String(),
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
