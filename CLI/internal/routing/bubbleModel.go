package routing

import (
	"fmt"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	keys     KeyMap
	list     list.Model
	choice   string
	styles   Styles
	quitting bool
	help     help.Model
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
		m.list.SetSize(msg.Width, msg.Height-5)
		return m, nil

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Enter):
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
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

	footer := m.styles.Help.Render("↑/↓/←/→: Navigate • Enter: Select • q/esc: Exit")

	v := tea.NewView("\n" + m.list.View() + "\n" + footer)
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

	const defaultWidth = 30
	const defaultHeight = 30

	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	m := Model{list: l, keys: createKeyMap(), help: help.New()}
	m.UpdateStyles(true)
	return m
}
