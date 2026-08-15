package types

type CountriesJSON []struct {
	Name      string    `json:"name"`
	Cioc      string    `json:"cioc,omitempty"`
	Latlng    []float32 `json:"latlng,omitempty"`
	Region    string    `json:"region"`
	Capital   string    `json:"capital,omitempty"`
	Languages []struct {
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"languages"`
	Subregion  string `json:"subregion"`
	Alpha2Code string `json:"alpha2Code"`
	Alpha3Code string `json:"alpha3Code"`
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies,omitempty"`
}
