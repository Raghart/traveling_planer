package bubbletea

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/utils"
)

type Model struct {
	keys     KeyMap
	list     list.Model
	styles   Styles
	quitting bool
	help     help.Model
	country  CountryManager
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

			return m, nil
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

	if m.quitting {
		return tea.NewView(m.styles.QuitText.Render("Hope to see you again!"))
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

func InitialModel() Model {
	items := []list.Item{
		Item{title: "Argentina", desc: "Land of tango, world-class beef, and the majestic Iguazu Falls."},
		Item{title: "Bolivia", desc: "Home to the otherworldly Salar de Uyuni and high-altitude culture."},
		Item{title: "Brazil", desc: "Vibrant rhythms of Rio, lush Amazon rainforest, and golden beaches."},
		Item{title: "Canada", desc: "Breathtaking Rocky Mountains and multicultural, friendly cities."},
		Item{title: "Chile", desc: "From the arid Atacama Desert to the glaciers of Patagonia."},
		Item{title: "Colombia", desc: "The world's finest coffee and the colorful streets of Cartagena."},
		Item{title: "Costa Rica", desc: "A tropical paradise for eco-tourism and 'Pura Vida' lifestyle."},
		Item{title: "Cuba", desc: "Time-capsule architecture, classic cars, and soul-stirring music."},
		Item{title: "Dominica", desc: "The 'Nature Island' with volcanic peaks and lush rainforests."},
		Item{title: "Dominican Republic", desc: "Pristine white-sand beaches and the historic Zona Colonial."},
		Item{title: "Grenada", desc: "The 'Spice Isle' famous for nutmeg, ginger, and aromatic air."},
		Item{title: "French Guiana", desc: "A unique blend of French culture and untamed Amazonian jungle."},
		Item{title: "Guyana", desc: "Untouched wilderness and the thundering Kaieteur Falls."},
		Item{title: "Saint Lucia", desc: "Iconic volcanic Pitons rising from the turquoise Caribbean Sea."},
		Item{title: "Honduras", desc: "Ancient Mayan ruins of Copan and world-class diving in Roatan."},
		Item{title: "Mexico", desc: "Delicious cuisine, ancient pyramids, and vibrant traditions."},
		Item{title: "Nicaragua", desc: "Land of lakes and volcanoes with charming colonial towns."},
		Item{title: "Panama", desc: "Where the canal meets the skyline and tropical islands."},
		Item{title: "Peru", desc: "Incan wonders of Machu Picchu and a world-renowned food scene."},
		Item{title: "Puerto Rico", desc: "Enchanting bioluminescent bays and the history of Old San Juan."},
		Item{title: "Paraguay", desc: "Authentic Guarani culture and the massive Itaipu Dam."},
		Item{title: "Suriname", desc: "A melting pot of cultures hidden in the tropical forest."},
		Item{title: "El Salvador", desc: "World-class surfing waves and beautiful volcanic landscapes."},
		Item{title: "Trinidad and Tobago", desc: "Home of Steelpan, Calypso, and an incredible Carnival."},
		Item{title: "United States", desc: "Iconic national parks, diverse cities, and endless road trips."},
		Item{title: "Venezuela", desc: "Angel Falls, the highest in the world, and Caribbean paradises."},
		Item{title: "Guatemala", desc: "Heart of the Mayan world and the stunning Lake Atitlan."},
		Item{title: "Belize", desc: "The Great Blue Hole and the longest barrier reef in the Americas."},
		Item{title: "Jamaica", desc: "Birthplace of Reggae, spicy jerk chicken, and island vibes."},
		Item{title: "Haiti", desc: "Rich history, resilient spirit, and vibrant, colorful art."},
		Item{title: "Bahamas", desc: "Crystal clear waters and over 700 stunning coral islands."},
		Item{title: "Barbados", desc: "Famous for its rum, cricket, and beautiful platinum coast."},
		Item{title: "Saint Kitts and Nevis", desc: "Historic fortresses and scenic railway journeys."},
		Item{title: "Antigua and Barbuda", desc: "365 distinct beaches—one for every day of the year."},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)

	m := Model{list: l, keys: createKeyMap(), help: help.New()}
	m.UpdateStyles(true)
	return m
}

func (m Model) GenerateProjectActions() []list.Item {
	return []list.Item{
		Item{
			title: "Temperature",
			desc:  fmt.Sprintf("Want to know the weekly temperature of %s?", m.country.To),
			task: func() {
				utils.PresentWeather(pubsub.GetTemperature(m.country.To))
			},
		},
		Item{
			title: "Currency",
			desc: fmt.Sprintf(
				"Want to know the currency difference from %s to %s?",
				m.country.From,
				m.country.To),
			task: func() {
				fmt.Printf("Currency value: %.2f", pubsub.GetCurrency(m.country.From, m.country.To))
			},
		},
	}
}
