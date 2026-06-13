package routing

type CountryManager struct {
	From              string
	Destination       string
	Description       string
	Currency          Currency
	ImageUrls         []string
	DailyTemperatures []DailyTemp
	Festivities       []FestivityData
}

type CountryData struct {
	DataType string
	Data     interface{}
	Error    error
}

type Currency struct {
	From         string
	FromCurrency string
	To           string
	ToCurrency   string
	ExchangeRate float32
	InverseRate  float32
}

type CountryDescription struct {
	Name        string
	Description string
}

type CountryAsciiImg struct {
	Name      string
	ImageUrls []string
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

type MsgError struct {
	Error error
}

type TestingData struct {
	IsCountry bool
}

type CountryFestivities struct {
	Country     string
	Festivities []FestivityData
}

type FestivityData struct {
	Date        string
	LocalName   string
	Name        string
	CountryCode string
	Fixed       bool
	Global      bool
	Types       []string
}
