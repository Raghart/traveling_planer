package bubbletea

import (
	"fmt"
	"io"

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

func (d itemDelegate) Height() int                                                     { return 1 }
func (d itemDelegate) Spacing() int                                                    { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd                         { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {}

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
