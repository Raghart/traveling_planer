package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Styles struct {
	Title        lipgloss.Style
	Item         lipgloss.Style
	SelectedItem lipgloss.Style
	Pagination   lipgloss.Style
	Help         lipgloss.Style
	QuitText     lipgloss.Style
}

func NewStyles(darkBG bool) Styles {
	var s Styles
	s.Title = lipgloss.NewStyle().MarginLeft(2)
	s.Item = lipgloss.NewStyle().PaddingLeft(4)
	s.SelectedItem = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	s.Pagination = list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4)
	s.Help = list.DefaultStyles(darkBG).HelpStyle.PaddingLeft(4).PaddingBottom(1)
	s.QuitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	return s
}

type Item string

func (i Item) FilterValue() string { return "" }

type ItemDelegate struct {
	Styles *Styles
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
			return d.Styles.SelectedItem.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Model struct {
	Styles   Styles
	List     list.Model
	Choice   string
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) UpdateStyles(isDark bool) {
	m.Styles = NewStyles(isDark)
	m.List.Styles.Title = m.Styles.Title
	m.List.Styles.PaginationStyle = m.Styles.Pagination
	m.List.Styles.HelpStyle = m.Styles.Help
	m.List.SetDelegate(ItemDelegate{Styles: &m.Styles})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyPressMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	if m.Choice != "" {
		return tea.NewView(m.Styles.QuitText.Render(fmt.Sprintf("Traveling from %s!", m.Choice)))
	}
	if m.Quitting {
		return tea.NewView(m.Styles.QuitText.Render("Have a good day!"))
	}
	return tea.NewView("\n" + m.List.View())
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

	s := NewStyles(true)

	l := list.New(items, &ItemDelegate{
		Styles: &s,
	}, defaultWidth, 14)
	l.Title = "Where are you From?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	m := Model{
		List: l,
	}
	return m
}

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
