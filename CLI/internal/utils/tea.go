package utils

import "github.com/Raghart/traveling_planer/internal/routing"

func InitialModel() routing.Model {
	return routing.Model{
		Choices: []string{
			"Argentina",
			"Bolivia",
			"Brazil",
			"Canada",
			"Chile",
			"Colombia",
			"Costa Rica",
			"Cuba",
			"Dominica",
			"Dominican Republic",
			"Grenada",
			"French Guiana",
			"Guyana",
			"Honduras",
			"Saint Lucia",
			"Mexico",
			"Nicaragua",
			"Panama",
			"Peru",
			"Puerto Rico",
			"Paraguay",
			"Suriname",
			"El Salvador",
			"Trinidad and Tobago",
			"United States",
			"Uruguay",
			"Venezuela",
			"Guatemala",
			"Belize",
			"Jamaica",
			"Haiti",
			"Bahamas",
			"Barbados",
			"Saint Kitts and Nevis",
			"Antigua and Barbuda",
		},
		Selected: make(map[int]struct{}),
	}
}
