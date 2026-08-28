package bubbletea

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Styles struct {
	Base            lipgloss.Style
	Red             color.Color
	Indigo          color.Color
	Green           color.Color
	HeaderText      lipgloss.Style
	StatusText      lipgloss.Style
	Status          lipgloss.Style
	StatusHeader    lipgloss.Style
	Highlight       lipgloss.Style
	ErrorHeaderText lipgloss.Style
	Help            lipgloss.Style
	StatusMessage   lipgloss.Style
}

func NewStyles(darkBG bool) *Styles {
	var s Styles
	lightDark := lipgloss.LightDark(darkBG)
	s.Red = lightDark(lipgloss.Color("#FE5F86"), lipgloss.Color("#FE5F86"))
	s.Indigo = lightDark(lipgloss.Color("#5A56E0"), lipgloss.Color("#5A56E0"))
	s.Green = lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BA84"))
	s.Base = lipgloss.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lipgloss.NewStyle().Foreground(s.Indigo).Bold(true).Padding(0, 1, 0, 2)
	s.StatusHeader = lipgloss.NewStyle().Foreground(s.Green).Bold(true)
	// Need to update StatusText to be presentable
	s.StatusText = lipgloss.NewStyle().Foreground(s.Indigo).Bold(true)

	s.Status = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(s.Indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.Highlight = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.Foreground(s.Red)
	s.Help = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lightDark(lipgloss.Color("#04B575"), lipgloss.Color("#04B575")))

	return &s
}

func (m *Model) AppFrameworkView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.Width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceStyle(lipgloss.NewStyle().Foreground(m.styles.Indigo)),
	)
}

func (m *Model) ErrorFrameworkView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.Width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceStyle(lipgloss.NewStyle().Foreground(m.styles.Red)),
	)
}

func (m *Model) ErrorView() string {
	var s string
	for _, err := range m.Form.Errors() {
		s += err.Error()
	}
	return s
}
