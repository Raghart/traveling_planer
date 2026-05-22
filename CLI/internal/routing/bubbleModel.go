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
	m.list.SetDelegate(list.NewDefaultDelegate())
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
				m.choice = i.title
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
		Item{title: "Argentina", desc: "The Best Carnita Asada meat in all South America!"},
		Item{title: "Bolivia", desc: "Land of extremes: soaring Andean peaks, vast salt flats, dense Amazon jungle, and rich indigenous traditions."},
		Item{title: "Brazil", desc: ""},
		Item{title: "Canada", desc: ""},
		Item{title: "Chile", desc: ""},
		Item{title: "Colombia", desc: ""},
		Item{title: "Costa Rica", desc: ""},
		Item{title: "Cuba", desc: ""},
		Item{title: "Dominica", desc: ""},
		Item{title: "Dominican Republic", desc: ""},
		Item{title: "Grenada", desc: ""},
		Item{title: "French Guiana", desc: ""},
		Item{title: "Guyana", desc: ""},
		Item{title: "Saint Lucia", desc: ""},
		Item{title: "Honduras", desc: ""},
		Item{title: "Mexico", desc: ""},
		Item{title: "Nicaragua", desc: ""},
		Item{title: "Panama", desc: ""},
		Item{title: "Peru", desc: ""},
		Item{title: "Puerto Rico", desc: ""},
		Item{title: "Paraguay", desc: ""},
		Item{title: "Suriname", desc: ""},
		Item{title: "El Salvador", desc: ""},
		Item{title: "Trinidad and Tobago", desc: ""},
		Item{title: "United States", desc: ""},
		Item{title: "Venezuela", desc: ""},
		Item{title: "Guatemala", desc: ""},
		Item{title: "Belize", desc: ""},
		Item{title: "Jamaica", desc: ""},
		Item{title: "Haiti", desc: ""},
		Item{title: "Bahamas", desc: ""},
		Item{title: "Barbados", desc: ""},
		Item{title: "Saint Kitts and Nevis", desc: ""},
		Item{title: "Antigua and Barbuda", desc: ""},
	}

	const defaultWidth = 30
	const defaultHeight = 30

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	m := Model{list: l, keys: createKeyMap(), help: help.New()}
	m.UpdateStyles(true)
	return m
}
