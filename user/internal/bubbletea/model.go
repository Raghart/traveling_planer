package bubbletea

import (
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/utils"
)

type Model struct {
	keys        KeyMap
	list        list.Model
	styles      Styles
	quitting    bool
	help        help.Model
	country     CountryManager
	viewport    viewport.Model
	renderer    *glamour.TermRenderer
	showResults bool
}

type CountryManager struct {
	From string
	To   string
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

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.RequestBackgroundColor,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-5)
		m.viewport.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height - 5)
		renderer, err := m.updateRenderer(msg.Width)
		if err != nil {
			log.Print(err)
		}
		m.renderer = renderer
		return m, nil

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Enter):
			i, ok := m.list.SelectedItem().(Item)
			if ok && m.country.From != "" && m.country.To == "" {
				m.country.To = i.title

				newItems := m.GenerateProjectActions()
				cmd := m.list.SetItems(newItems)

				m.list.Title = fmt.Sprintf("Traveling from %s to %s", m.country.From, m.country.To)
				m.list.NewStatusMessage("")
				return m, cmd
			}

			if ok && m.country.From == "" {
				m.country.From = i.title
				m.list.Title = "Where do you want to go?"
				m.list.NewStatusMessage(
					m.styles.StatusMessage.Render("From: " + i.title))
				m.list.RemoveItem(m.list.Index())
			}

			if ok && m.country.From != "" && m.country.To != "" {
				formattedMsg := i.task()
				str, _ := m.renderer.Render(formattedMsg)
				m.viewport.SetContent(str)
				m.showResults = true
			}

			return m, nil
		case key.Matches(msg, m.keys.Quit):
			if m.showResults {
				m.showResults = false
				return m, nil
			}

			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {

	if m.quitting {
		return tea.NewView(m.styles.QuitText.Render("Hope to see you again!"))
	}

	if m.showResults {
		return tea.NewView(m.viewport.View())
	}

	footer := m.styles.Help.Render(
		strings.Join([]string{
			"↑/↓/←/→: Navigate",
			"Enter: Select",
			"Filter: /",
			"q/esc: Exit",
		}, " • "))

	strView := strings.Join([]string{
		m.list.View(),
		footer,
	}, "\n")

	v := tea.NewView(strView)
	v.AltScreen = true
	return v
}

func (m Model) GenerateProjectActions() []list.Item {
	return []list.Item{
		Item{
			title: "Temperature",
			desc:  fmt.Sprintf("Want to know the weekly temperature of %s?", m.country.To),
			task: func() string {
				return utils.PresentWeather(pubsub.GetTemperature(m.country.To))
			},
		},
		Item{
			title: "Currency",
			desc: fmt.Sprintf(
				"Want to know the currency difference from %s to %s?",
				m.country.From,
				m.country.To),
			task: func() string {
				return utils.FormatCurrency(pubsub.GetCurrency(m.country.From, m.country.To))
			},
		},
	}
}

func (m Model) updateRenderer(terminalWidth int) (*glamour.TermRenderer, error) {
	glamourRenderWidth := terminalWidth - m.viewport.Style.GetHorizontalFrameSize() - 3
	styles := styles.DarkStyleConfig

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styles),
		glamour.WithWordWrap(glamourRenderWidth),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to create the renderer: %w", err)
	}
	return renderer, nil
}

func InitialModel() Model {
	items := GenerateCountryItems()
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)

	vp := CreateViewportRenderer()

	m := Model{
		list:     l,
		keys:     createKeyMap(),
		help:     help.New(),
		viewport: vp,
	}
	m.UpdateStyles(true)
	return m
}
