package utils

import (
	"charm.land/bubbles/v2/key"
	"charm.land/lipgloss/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
)

func NewStyles(darkBG bool) routing.Styles {
	lightDark := lipgloss.LightDark(darkBG)
	return routing.Styles{
		App: lipgloss.NewStyle().Padding(1, 2),
		Title: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1),
		StatusMessage: lipgloss.NewStyle().
			Foreground(lightDark(lipgloss.Color("#04B575"), lipgloss.Color("#04B575"))),
	}
}

type Item struct {
	Title       string
	Description string
}

func (i Item) GiveTitle() string       { return i.Title }
func (i Item) GiveDescription() string { return i.Description }
func (i Item) FilterValue() string     { return i.Title }

type ListKeyMap struct {
	ToggleSpinner    key.Binding
	ToggleTitleBar   key.Binding
	ToggleStatusBar  key.Binding
	TogglePagination key.Binding
	ToggleHelpMenu   key.Binding
	InsertItem       key.Binding
}

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
