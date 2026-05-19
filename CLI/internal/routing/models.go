package routing

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Currency struct {
	From  string
	To    string
	Value float32
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

type CountryData struct {
	IsCountry bool
}

type Styles struct {
	Title      lipgloss.Style
	Item       lipgloss.Style
	SelectItem lipgloss.Style
	Pagination lipgloss.Style
	Help       lipgloss.Style
	QuitText   lipgloss.Style
}

type Model struct {
	Styles   Styles
	List     list.Model
	Choice   string
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}
