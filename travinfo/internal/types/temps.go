package types

type CountryLocation struct {
	Latitude  float64
	Longitude float64
}

type CountryTemp struct {
	Country string
	Value   float32
}
