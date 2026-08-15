package bubbletea

import (
	"fmt"
	"io"
	"log"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/utils"
)

type Item struct {
	title, desc string
	task        func() string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type itemDelegate struct {
	styles *Styles
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {

}

func GenerateCountryItems() []list.Item {
	return []list.Item{
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
}

func GenerateFromCountryItems() []list.Item {
	countryList, err := utils.GetAllCountriesData()
	if err != nil {
		log.Fatal(err)
	}

	countryItems := []list.Item{}
	for _, countryData := range countryList {
		countryItems = append(countryItems, Item{title: countryData.Name, desc: ""})
	}

	return countryItems
}

func (m *Model) GenerateProjectActions() []list.Item {
	return []list.Item{
		Item{
			title: "Temperature",
			desc:  fmt.Sprintf("Want to know the weekly temperature of %s?", m.country.Destination),
			task: func() string {
				return utils.FormatWeather(m.country.DailyTemperatures)
			},
		},
		Item{
			title: "Currency",
			desc: fmt.Sprintf(
				"Want to know the currency difference from %s to %s?",
				m.country.From,
				m.country.Destination),
			task: func() string {
				return utils.FormatCurrency(m.country.Currency)
			},
		},
		Item{
			title: "Holidays",
			desc: fmt.Sprintf(
				"Want to know the holidays of %s to plan your adventure?", m.country.Destination),
			task: func() string {
				return utils.FormatHolidays(m.country.Destination, m.country.Festivities)
			},
		},
		Item{
			title: "Description",
			desc: fmt.Sprintf(
				"Want to read more about %s?", m.country.Destination),
			task: func() string {
				return utils.FormatDescription(m.country.Destination, m.country.Description)
			},
		},
		Item{
			title: fmt.Sprintf("%s's Images", m.country.Destination),
			desc:  fmt.Sprintf("Check out images of %s!", m.country.Destination),
			task: func() string {
				return utils.FormatImageUrls(m.country.Destination, m.country.ImageUrls)
			},
		},
	}
}
