package types

type CountryLocation struct {
	Latitude  float64
	Longitude float64
}

type CountryTemp struct {
	Country           string
	DailyTemperatures []DailyTemp
	Value             float64
}

type DailyTemp struct {
	WeatherCode int
	Max         float64
	Min         float64
	RainProb    int
	AparentMax  float64
	AparentMin  float64
}

type CountryTempJSON struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	DailyUnits           struct {
		Time                        string `json:"time"`
		Temperature2MMax            string `json:"temperature_2m_max"`
		Temperature2MMin            string `json:"temperature_2m_min"`
		WeatherCode                 string `json:"weather_code"`
		PrecipitationProbabilityMax string `json:"precipitation_probability_max"`
		ApparentTemperatureMax      string `json:"apparent_temperature_max"`
		ApparentTemperatureMin      string `json:"apparent_temperature_min"`
	} `json:"daily_units"`
	Daily struct {
		Time                        []string  `json:"time"`
		Temperature2MMax            []float64 `json:"temperature_2m_max"`
		Temperature2MMin            []float64 `json:"temperature_2m_min"`
		WeatherCode                 []int     `json:"weather_code"`
		PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
		ApparentTemperatureMax      []float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin      []float64 `json:"apparent_temperature_min"`
	} `json:"daily"`
}
