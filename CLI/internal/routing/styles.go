package routing

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
}

func NewStyles(darkBG bool) Styles {
	var s Styles
	lightDark := lipgloss.LightDark(darkBG)
	s.App = lipgloss.NewStyle().Padding(1, 2)
	s.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("#25A065")).
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	s.Item = lipgloss.NewStyle().PaddingLeft(4)
	s.SelectedItem = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	s.Pagination = list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4)
	s.Help = list.DefaultStyles(darkBG).HelpStyle.PaddingLeft(4).PaddingBottom(1)
	s.QuitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lightDark(lipgloss.Color("#04B575"), lipgloss.Color("#04B575")))
	return s
}
