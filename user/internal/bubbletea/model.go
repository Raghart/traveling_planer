package bubbletea

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
)

type Model struct {
	keys        KeyMap
	list        list.Model
	styles      *Styles
	help        help.Model
	country     routing.CountryManager
	progress    progress.Model
	Form        *huh.Form
	Width       int
	debugStrs   []string
	loadingData bool
	dataCh      chan tea.Msg
	dataMsg     string
}

func (m *Model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = min(msg.Width, 80) - m.styles.Base.GetHorizontalFrameSize()
		m.list.SetSize(msg.Width, msg.Height-5)
		return m, nil

	case progress.FrameMsg:
		m.progress, cmd = m.progress.Update(msg)
		return m, cmd

	case routing.CountryData:
		switch msg.DataType {
		case "temp":
			tempSlice, _ := msg.Data.([]routing.DailyTemp)
			m.country.DailyTemperatures = tempSlice
			m.dataMsg = fmt.Sprintf("%s temperature loaded (+20%%)", m.country.Destination)
		case "currency":
			currencyData, _ := msg.Data.(routing.Currency)
			m.country.Currency = currencyData
			m.dataMsg = fmt.Sprintf("%s currency loaded (+20%%)", m.country.Destination)
		case "holidays":
			holidaysData, _ := msg.Data.([]routing.FestivityData)
			m.country.Festivities = holidaysData
			m.dataMsg = fmt.Sprintf("%s holidays loaded (+20%%)", m.country.Destination)
		case "description":
			description, _ := msg.Data.(string)
			m.country.Description = description
			m.dataMsg = fmt.Sprintf("%s description loaded (+20%%)", m.country.Destination)
		case "images":
			urlImages, _ := msg.Data.([]string)
			m.country.ImageUrls = urlImages
			m.dataMsg = fmt.Sprintf("%s url images loaded (+20%%)", m.country.Destination)
		}

		return m, tea.Batch(
			m.progress.IncrPercent(.20),
			ListenCountryData(m.dataCh),
			m.CheckPercentage(),
		)

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() tea.View {
	errorSlice := m.Form.Errors()
	header := m.AppFrameworkView("Traveling Planner")

	if len(errorSlice) > 0 {
		header = m.ErrorFrameworkView(m.ErrorView())
	}

	v := strings.TrimSuffix(m.Form.View(), "\n\n")
	form := lipgloss.NewStyle().Margin(1, 0).Render(v)

	body := lipgloss.JoinHorizontal(lipgloss.Left, form)

	footer := m.AppFrameworkView(m.Form.Help().ShortHelpView(m.Form.KeyBinds()))
	if len(errorSlice) > 0 {
		footer = m.ErrorFrameworkView("")
	}

	if m.loadingData {
		return tea.NewView("\n" +
			fmt.Sprintf("Loading data from %s...\n\n", m.country.Destination) +
			fmt.Sprintf("%s\n\n", m.dataMsg) + m.progress.View() + "\n\n" +
			HelpStyle("Press 'q' to quit"))
	}

	//if len(m.debugStrs) > 1 {
	//	strView += strings.Join(m.debugStrs, "\n")
	//}

	return tea.NewView(m.styles.Base.Render(header + "\n" + body + "\n\n" + footer))
}

func InitialModel() *Model {
	countryNames, err := utils.GetAllCountriesNames()
	if err != nil {
		log.Fatal(err)
	}

	currentDate := time.Now().Format(time.DateOnly)
	var budget string

	newForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("country").
				Options(huh.NewOptions(countryNames...)...).
				Title("Where are you from?").
				Height(10).
				Description("Select your origin's country"),

			huh.NewSelect[string]().
				Key("budget").
				Options(huh.NewOptions("Economic", "Moderate", "High Level")...).
				Title("How big is your budget?").
				Value(&budget).
				DescriptionFunc(func() string {
					switch budget {
					case "Economic":
						return "Budget equal or lower than 5.000$"
					case "Moderate":
						return "Budget Between 5.000$ and 10.000$"
					case "High Level":
						return "Budget highter than 10.000$"
					default:
						return "Select your budget"
					}
				}, &budget),

			huh.NewInput().
				Key("travelDate").
				Title("When are you planning to travel?").
				CharLimit(10).
				Description("Date Format YYYY-MM-DD | YYYY-M-DD").
				Placeholder("YYYY-MM-DD").
				Value(&currentDate).
				Validate(func(s string) error {
					rawDateSlice := strings.Split(s, "-")
					if len(rawDateSlice) != 3 {
						return errors.New("Error: Invalid formtat, date must be YYYY-MM-DD")
					}

					parsedDateSlice := []string{}
					for idx, dateStr := range rawDateSlice {
						dateNum, err := strconv.ParseInt(dateStr, 10, 32)
						if err != nil {
							return fmt.Errorf("Error: %v is not a valid number!", dateStr)
						}

						switch idx {
						case 0:
							if dateNum < int64(time.Now().Year()) {
								return errors.New("Error: Invalid year, It is impossible to travel in the past!")
							}
							if dateNum > 3000 {
								return errors.New("Error: Year is too far into the future!")
							}
						case 1:
							if dateNum < 1 || dateNum > 12 {
								return errors.New("Error: Invalid month, it should be a number between 1 and 12!")
							}
						case 2:
							if dateNum < 1 || dateNum > 31 {
								return errors.New("Error: Invalid day of the month!")
							}
						}
						parsedDateSlice = append(parsedDateSlice, fmt.Sprintf("%02d", dateNum))
					}

					rawTime, err := time.Parse(time.DateOnly, strings.Join(parsedDateSlice, "-"))
					if err != nil {
						return fmt.Errorf("Error: Invalid time: %v", err)
					}

					dateNow, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
					if rawTime.Before(dateNow) {
						return errors.New("Invalid date: Traveling date can't be in the past!")
					}

					return nil
				}),
		),
		huh.NewGroup(
			huh.NewInput().
				Key("travelDuration").
				Title("How many days are you planning to travel?").
				CharLimit(3).
				Placeholder("7").
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 32)
					if err != nil {
						return fmt.Errorf("Error: number is not a valid int: %v", err)
					}

					return nil
				}).
				Description("Select the number of days you're planning to travel"),

			huh.NewConfirm().
				Key("enviroment").
				Title("Do you prefer hotter or colder enviroments?").
				Affirmative("Hotter").
				Negative("Colder"),

			huh.NewMultiSelect[string]().
				Key("activities").
				Options(
					huh.NewOption("Nature-wilderness", "nature"),
					huh.NewOption("City-life style", "city"),
					huh.NewOption("History-culture", "history"),
				).
				Title("What do you enjoy doing while traveling?").
				Description("Select the activities you enjoy doing!").
				Height(4),
		),
	).
		WithShowHelp(false).
		WithShowErrors(false)

	m := &Model{
		keys:     createKeyMap(),
		help:     help.New(),
		Form:     newForm,
		progress: progress.New(progress.WithDefaultBlend()),
		dataMsg:  "Downloading required data...",
	}

	return m
}
