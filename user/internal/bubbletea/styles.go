package bubbletea

import (
	"image/color"

	"charm.land/bubbles/v2/list"
	"charm.land/lipgloss/v2"
)

type Styles struct {
	App           lipgloss.Style
	Title         lipgloss.Style
	Item          lipgloss.Style
	SelectedItem  lipgloss.Style
	Pagination    lipgloss.Style
	Help          lipgloss.Style
	QuitText      lipgloss.Style
	StatusMessage lipgloss.Style
	ActiveDot     lipgloss.Style
	InactiveDot   lipgloss.Style
	LightPurple   lipgloss.Style

	Base            lipgloss.Style
	Red             color.Color
	Indigo          color.Color
	Green           color.Color
	HeaderText      lipgloss.Style
	Status          lipgloss.Style
	StatusHeader    lipgloss.Style
	Highlight       lipgloss.Style
	ErrorHeaderText lipgloss.Style
	HelpForm        lipgloss.Style
}

func NewStyles(darkBG bool) *Styles {
	var s Styles
	lightDark := lipgloss.LightDark(darkBG)

	s.Red = lightDark(lipgloss.Color("#FE5F86"), lipgloss.Color("#FE5F86"))
	s.Indigo = lightDark(lipgloss.Color("#5A56E0"), lipgloss.Color("#5A56E0"))
	s.Green = lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BA84"))
	s.Base = lipgloss.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lipgloss.NewStyle().Foreground(s.Indigo).Bold(true).Padding(0, 1, 0, 2)
	s.Status = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(s.Indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lipgloss.NewStyle().Foreground(s.Green).Bold(true)
	s.Highlight = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.Foreground(s.Red)
	s.HelpForm = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	s.App = lipgloss.NewStyle().Padding(1, 2)
	s.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFF")).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("#25A065")).
		Padding(0, 1).
		Bold(true)

	s.Item = lipgloss.NewStyle().PaddingLeft(4)
	s.SelectedItem = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	s.Pagination = list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4)
	s.QuitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lightDark(lipgloss.Color("#04B575"), lipgloss.Color("#04B575")))

	s.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("#12d142"))
	s.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	s.LightPurple = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return &s
}

func HelpStyle(text string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render(text)
}

func helpViewport() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("\n ↑/↓: Navigate • q/esc: Quit\n")
}

func (m *Model) UpdateStyles(isDark bool) {
	m.styles = NewStyles(isDark)
	m.list.Styles.Title = m.styles.Title
	m.list.Styles.PaginationStyle = m.styles.Pagination
	m.list.Styles.HelpStyle = m.styles.Help

	m.list.Styles.ActivePaginationDot = m.styles.ActiveDot
	m.list.Styles.InactivePaginationDot = m.styles.InactiveDot
	m.list.Paginator.ActiveDot = m.styles.ActiveDot.Render("•")
	m.list.Paginator.InactiveDot = m.styles.InactiveDot.Render("•")
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
