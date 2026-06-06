package routing

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

const (
	ExchangePerilDirect = "peril_direct"
	TestingKey          = "pause"
)
