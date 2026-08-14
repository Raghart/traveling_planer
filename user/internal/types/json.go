package types

type CountriesJSON []struct {
	Area int    `json:"area,omitempty"`
	Cioc string `json:"cioc,omitempty"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps,omitempty"`
	Name  string `json:"name"`
	Flags struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
	Latlng    []int    `json:"latlng,omitempty"`
	Region    string   `json:"region"`
	Borders   []string `json:"borders,omitempty"`
	Capital   string   `json:"capital,omitempty"`
	Demonym   string   `json:"demonym"`
	Languages []struct {
		Name       string `json:"name"`
		Iso6391    string `json:"iso639_1"`
		Iso6392    string `json:"iso639_2"`
		NativeName string `json:"nativeName"`
	} `json:"languages"`
	Subregion  string   `json:"subregion"`
	Timezones  []string `json:"timezones"`
	Alpha2Code string   `json:"alpha2Code"`
	Alpha3Code string   `json:"alpha3Code"`
	Currencies []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies,omitempty"`
	NativeName   string   `json:"nativeName"`
	Population   int      `json:"population"`
	Independent  bool     `json:"independent"`
	NumericCode  string   `json:"numericCode"`
	AltSpellings []string `json:"altSpellings,omitempty"`
	CallingCodes []string `json:"callingCodes"`
	Translations struct {
		Br string `json:"br"`
		De string `json:"de"`
		Es string `json:"es"`
		Fa string `json:"fa"`
		Fr string `json:"fr"`
		Hr string `json:"hr"`
		Hu string `json:"hu"`
		It string `json:"it"`
		Ja string `json:"ja"`
		Nl string `json:"nl"`
		Pt string `json:"pt"`
	} `json:"translations"`
	RegionalBlocs []struct {
		Name    string `json:"name"`
		Acronym string `json:"acronym"`
	} `json:"regionalBlocs,omitempty"`
	TopLevelDomain    []string `json:"topLevelDomain"`
	PopulationDensity float64  `json:"populationDensity,omitempty"`
	Gini              float64  `json:"gini,omitempty"`
}
