package routing

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	keys     KeyMap
	list     list.Model
	choice   string
	styles   Styles
	quitting bool
}

func (m *Model) UpdateStyles(isDark bool) {
	m.styles = NewStyles(isDark)
	m.list.Styles.Title = m.styles.Title
	m.list.Styles.PaginationStyle = m.styles.Pagination
	m.list.Styles.HelpStyle = m.styles.Help
	m.list.SetDelegate(itemDelegate{styles: &m.styles})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.RequestBackgroundColor,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyPressMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	if m.choice != "" {
		return tea.NewView(m.styles.QuitText.Render(fmt.Sprintf("%s, Understood!", m.choice)))
	}
	if m.quitting {
		return tea.NewView(m.styles.QuitText.Render("Hope to see you again!"))
	}

	v := tea.NewView("\n" + m.list.View())
	v.AltScreen = true
	return v
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
