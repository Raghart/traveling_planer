package routing

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
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
	return nil
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
	return tea.NewView("\n" + m.list.View())
}
