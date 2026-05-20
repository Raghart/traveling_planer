package routing

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Item string

func (i Item) FilterValue() string { return "" }

type itemDelegate struct {
	styles *Styles
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := d.styles.Item.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.styles.SelectedItem.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func InitialModel() Model {
	items := []list.Item{
		Item("Argentina"),
		Item("Bolivia"),
		Item("Brazil"),
		Item("Canada"),
		Item("Chile"),
		Item("Colombia"),
		Item("Costa Rica"),
		Item("Cuba"),
		Item("Dominica"),
		Item("Dominican Republic"),
		Item("Grenada"),
		Item("French Guiana"),
		Item("Guyana"),
		Item("Saint Lucia"),
		Item("Honduras"),
		Item("Mexico"),
		Item("Nicaragua"),
		Item("Panama"),
		Item("Peru"),
		Item("Puerto Rico"),
		Item("Paraguay"),
		Item("Suriname"),
		Item("El Salvador"),
		Item("Trinidad and Tobago"),
		Item("United States"),
		Item("Uruguay"),
		Item("Venezuela"),
		Item("Guatemala"),
		Item("Belize"),
		Item("Jamaica"),
		Item("Haiti"),
		Item("Bahamas"),
		Item("Barbados"),
		Item("Saint Kitts and Nevis"),
		Item("Antigua and Barbuda"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, 14)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	m := Model{list: l}
	m.UpdateStyles(true)
	return m
}
