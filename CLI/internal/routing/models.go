package routing

import (
	"fmt"

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
	App           lipgloss.Style
	Title         lipgloss.Style
	StatusMessage lipgloss.Style
}

type Model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter", "space":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	s := "Where are you from?\n\n"
	for idx, choice := range m.Choices {
		cursor := " "
		if m.Cursor == idx {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[idx]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit.\n"
	return tea.NewView(s)
}
