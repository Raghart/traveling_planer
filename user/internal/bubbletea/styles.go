package bubbletea

import (
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
}

func NewStyles(darkBG bool) Styles {
	var s Styles
	lightDark := lipgloss.LightDark(darkBG)
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
	s.Help = list.DefaultStyles(darkBG).HelpStyle.PaddingLeft(4).PaddingBottom(1).
		Foreground(lipgloss.Color("241"))
	s.QuitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lightDark(lipgloss.Color("#04B575"), lipgloss.Color("#04B575")))

	s.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("#12d142"))
	s.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	return s
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

func (m *Model) UpdateDateStyles() {
	tiStyles := m.TextInput.Styles()
	tiStyles.Cursor.Color = lipgloss.Color("205")
	tiStyles.Focused.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	tiStyles.Focused.Text = lipgloss.NewStyle().Foreground(lipgloss.White)
	tiStyles.Blurred.Prompt = lipgloss.NewStyle().Foreground(lipgloss.White)
	tiStyles.Focused.Text = lipgloss.NewStyle().Foreground(lipgloss.White)
	m.TextInput.SetStyles(tiStyles)
}
