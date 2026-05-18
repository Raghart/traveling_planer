package utils

import "github.com/Raghart/traveling_planer/internal/routing"

func InitialModel() routing.Model {
	return routing.Model{
		Choices: []string{
			"",
		},
		Selected: make(map[int]struct{}),
	}
}
