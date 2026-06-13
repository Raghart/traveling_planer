package utils

import (
	"fmt"
	"time"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
)

func LoadCurrencyMap() map[string]types.CurrType {
	return map[string]types.CurrType{
		"Argentina":             types.ArgentinaPeso,
		"Bolivia":               types.BolivianCurrency,
		"Brazil":                types.BrazilianCurrency,
		"Canada":                types.CanadianDollars,
		"Chile":                 types.ChileCurrency,
		"Colombia":              types.ColombianPeso,
		"Costa Rica":            types.CostaCurrency,
		"Cuba":                  types.CubaCurrency,
		"Dominica":              types.DominicanCurrency,
		"Dominican Republic":    types.DominicanCurrency,
		"Ecuador":               types.Dollars,
		"Grenada":               types.AntiguaDollars,
		"French Guiana":         types.Euros,
		"Guyana":                types.GuyanaDollars,
		"Honduras":              types.HondurasCurrency,
		"Saint Lucia":           types.AntiguaDollars,
		"Mexico":                types.MexicanPeso,
		"Nicaragua":             types.NicaraguaCurrency,
		"Panama":                types.Dollars,
		"Peru":                  types.PeruvianCurrency,
		"Puerto Rico":           types.Dollars,
		"Paraguay":              types.ParaguayCurrency,
		"Suriname":              types.SurinameDollars,
		"El Salvador":           types.Dollars,
		"Trinidad and Tobago":   types.TrinidadDollars,
		"United States":         types.Dollars,
		"Uruguay":               types.UruguayCurrency,
		"Venezuela":             types.Dollars,
		"Guatemala":             types.GuatemaQuetzal,
		"Belize":                types.BelizeDollars,
		"Jamaica":               types.JamaicanDollars,
		"Haiti":                 types.Dollars,
		"Bahamas":               types.BahamasDollars,
		"Barbados":              types.BarbadosDollars,
		"Saint Kitts and Nevis": types.AntiguaDollars,
		"Antigua and Barbuda":   types.AntiguaDollars,
	}
}

func LoadCountryCodesMap() map[string]string {
	return map[string]string{
		"Argentina":             "AR",
		"Bolivia":               "BO",
		"Brazil":                "BR",
		"Canada":                "CA",
		"Chile":                 "CL",
		"Colombia":              "CO",
		"Costa Rica":            "CR",
		"Cuba":                  "CU",
		"Dominica":              "DM",
		"Dominican Republic":    "DO",
		"Ecuador":               "EC",
		"Grenada":               "GD",
		"French Guiana":         "GF",
		"Guyana":                "GY",
		"Honduras":              "HN",
		"Saint Lucia":           "LC",
		"Mexico":                "MX",
		"Nicaragua":             "NI",
		"Panama":                "PA",
		"Peru":                  "PE",
		"Puerto Rico":           "PR",
		"Paraguay":              "PY",
		"Suriname":              "SR",
		"El Salvador":           "SV",
		"Trinidad and Tobago":   "TT",
		"United States":         "US",
		"Uruguay":               "UY",
		"Venezuela":             "VE",
		"Guatemala":             "GT",
		"Belize":                "BZ",
		"Jamaica":               "JM",
		"Haiti":                 "HT",
		"Bahamas":               "BS",
		"Barbados":              "BB",
		"Saint Kitts and Nevis": "KN",
		"Antigua and Barbuda":   "AG",
	}
}

func LoadSpecialCountryFestivities() map[string][]types.FestivityData {
	actualYear := time.Now().Year()
	return map[string][]types.FestivityData{
		"Saint Kitts and Nevis": []types.FestivityData{},
		"Antigua and Barbuda":   []types.FestivityData{},
		"Trinidad and Tobago": {
			{
				Name: "Trinidad & Tobago Carnival",
				Date: fmt.Sprintf("February 16th to 17th, %d", actualYear),
			},
			{
				Name: "Buccoo goat & Crab race festival",
				Date: fmt.Sprintf("April 7th, %d", actualYear),
			},
			{
				Name: "Tobago Jazz, Music & Golf Weekend",
				Date: fmt.Sprintf("May 1st to 3rd, %d", actualYear),
			},
			{
				Name: "Tobago Heritage Festival",
				Date: fmt.Sprintf("July 1st to August 1st %d", actualYear),
			},
			{
				Name: "Great race",
				Date: fmt.Sprintf("August 15th %d", actualYear),
			},
			{
				Name: "Tobago blue Food Festival",
				Date: fmt.Sprintf("October 18th %d", actualYear),
			},
			{
				Name: "Tobago Carnival",
				Date: fmt.Sprintf("October 30th - November 1st %d", actualYear),
			},
		},
		"Saint Lucia": {
			{
				Name: "Saint Lucia Jazz & Arts Festival",
				Date: fmt.Sprintf("April 30th - May 9th, %d", actualYear),
			},
			{
				Name: "Saint Lucia Carnival",
				Date: fmt.Sprintf("July 1st - 22nd, %d", actualYear),
			},
			{
				Name: "Mercury Fest Weekend",
				Date: fmt.Sprintf("August 14th - 16th, %d", actualYear),
			},
			{
				Name: "La Rose Flower Festival",
				Date: fmt.Sprintf("August 30th, %d", actualYear),
			},
			{
				Name: "Creole Heritage Month",
				Date: fmt.Sprintf("October %d", actualYear),
			},
			{
				Name: "La Marguerite Flower Festival",
				Date: fmt.Sprintf("October 17, %d", actualYear),
			},
			{
				Name: "National Day - Festival of Lights & Renewal",
				Date: fmt.Sprintf("December 13, %d", actualYear),
			},
		},
		"French Guiana": []types.FestivityData{},
		"Dominica":      []types.FestivityData{},
	}
}
