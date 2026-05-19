package utils

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
)

func NewStyles(darkBG bool) routing.Styles {
	var s routing.Styles
	s.Title = lipgloss.NewStyle().MarginLeft(2)
	s.Item = lipgloss.NewStyle().PaddingLeft(4)
	s.SelectItem = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	s.Pagination = list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4)
	s.Help = list.DefaultStyles(darkBG).HelpStyle.PaddingLeft(4).PaddingBottom(1)
	s.QuitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	return s
}

type Item string

func (i Item) FilterValue() string { return "" }

type ItemDelegate struct {
	Styles *routing.Styles
}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i)
	fn := d.Styles.Item.Render

	if index == m.Index() {
		fn = func(s ...string) string {
			return d.Styles.SelectItem.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

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
