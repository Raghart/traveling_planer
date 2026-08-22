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
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/glamour/v2"
	"charm.land/huh/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
	"github.com/Raghart/traveling_planer/internal/utils"
)

type Model struct {
	keys          KeyMap
	list          list.Model
	styles        Styles
	help          help.Model
	country       routing.CountryManager
	progress      progress.Model
	viewport      viewport.Model
	Form          *huh.Form
	renderer      *glamour.TermRenderer
	questionTitle string
	isWritingDate bool
	TextInput     textinput.Model
	debugStrs     []string
	showResults   bool
	quitting      bool
	loadingData   bool
	dataCh        chan tea.Msg
	dataMsg       string
	err           error
}

func (m *Model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	return m, cmd
	/*
		var cmd tea.Cmd
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetSize(msg.Width, msg.Height-5)
			m.viewport.SetWidth(msg.Width)
			m.viewport.SetHeight(msg.Height - 8)
			renderer, err := m.updateRenderer(msg.Width)
			if err != nil {
				m.err = err
				return m, nil
			}

			m.renderer = renderer
			return m, nil

		case progress.FrameMsg:
			m.progress, cmd = m.progress.Update(msg)
			return m, cmd

		case routing.CountryData:
			if msg.Error != nil {
				m.err = msg.Error
				return m, nil
			}

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
			case key.Matches(msg, m.keys.Help):
				m.help.ShowAll = !m.help.ShowAll

			case key.Matches(msg, m.keys.Enter):
				i, ok := m.list.SelectedItem().(Item)

				if ok && m.country.From == "" {
					m.country.From = i.title
					m.country.UserData.From = i.title

					m.list.Title = "How big is your budget?"
					m.list.NewStatusMessage(
						m.styles.StatusMessage.Render("From: " + i.title))

					m.list.SetItems(GenerateBudgetItems())
					m.list.ResetFilter()
					return m, nil
				}

				if ok && m.country.UserData.Budget == "" {
					m.country.UserData.Budget = i.title
					m.list.NewStatusMessage(m.styles.StatusMessage.Render(
						fmt.Sprintf("%s budget registered!", i.title),
					))

					m.isWritingDate = true
					return m, nil
				}

				if m.country.UserData.TravelDate.IsZero() {
					rawTravelDate := m.TextInput.Value()
					parsedTravelDate, err := utils.ParseRawDateStr(rawTravelDate)
					if err != nil {
						m.debugStrs = append(m.debugStrs, fmt.Sprintf("Error: %v", err))
						return m, nil
					}

					userTravelDate, err := time.Parse(time.DateOnly, parsedTravelDate)

					if err != nil {
						m.debugStrs = append(m.debugStrs, fmt.Sprintf("ERROR: %v", err))
						return m, nil
					}

					dateNow, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
					if userTravelDate.Before(dateNow) {
						m.debugStrs = append(m.debugStrs, "Error: Travel Date can't be in the past")
						return m, nil
					}

					m.country.UserData.TravelDate = userTravelDate
					m.questionTitle = "How long will your travel be?"
					m.TextInput.SetValue(userTravelDate.Add(24 * time.Hour).Format(time.DateOnly))
					return m, nil
				}

				if m.country.UserData.TravelEndDate.IsZero() {
					rawUserTravelEnds := m.TextInput.Value()
					parsedTravelEnds, err := utils.ParseRawDateStr(rawUserTravelEnds)
					if err != nil {
						m.debugStrs = append(m.debugStrs, fmt.Sprintf("Error: %v", err))
						return m, nil
					}

					userTravelEnds, err := time.Parse(time.DateOnly, parsedTravelEnds)
					if err != nil {
						m.debugStrs = append(m.debugStrs,
							fmt.Sprintf("Error: unable to parse that time: %v", err))
						return m, nil
					}

					if userTravelEnds.Before(m.country.UserData.TravelDate) {
						m.debugStrs = append(
							m.debugStrs, "Error: Travel Length can't be before the traveling start!")
					}
					m.country.UserData.TravelEndDate = userTravelEnds
					return m, nil
				}

				return m, nil

			case key.Matches(msg, m.keys.Quit):
				if m.showResults {
					m.showResults = false
					m.loadingData = false
					return m, nil
				}

				if m.err != nil {
					m.err = nil
					return m, nil
				}

				m.quitting = true
				return m, tea.Quit

			case m.isWritingDate:
				allowedKeys := []string{
					"left", "right", "backspace", "-",
					"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
				}

				switch {
				case slices.Contains(allowedKeys, msg.Keystroke()):
					m.TextInput, cmd = m.TextInput.Update(msg)
					return m, cmd
				}
			case m.list.FilterState() == list.Filtering:
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}
		}

		m.list, cmd = m.list.Update(msg)
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	*/
}

func (m *Model) View() tea.View {
	return tea.NewView(m.Form.View())
	/*
		if m.err != nil {
			return tea.NewView(m.styles.QuitText.Render(
				fmt.Sprintf("%v\n\n", m.err) +
					HelpStyle("Press 'q' to exit error window")))
		}

		if m.loadingData {
			return tea.NewView("\n" +
				fmt.Sprintf("Loading data from %s...\n\n", m.country.Destination) +
				fmt.Sprintf("%s\n\n", m.dataMsg) +
				m.progress.View() + "\n\n" + HelpStyle("Press 'q' to quit"))
		}

		if m.quitting {
			return tea.NewView(m.styles.QuitText.Render("Hope to see you again!"))
		}

		if m.showResults {
			return tea.NewView(m.viewport.View() + helpViewport())
		}

		footer := m.styles.Help.Render(m.buildDynamicHelp())

		var strView string
		var c *tea.Cursor

		if m.isWritingDate {
			headerView := m.styles.Title.Render(m.questionTitle)
			dateFormat := m.styles.LightPurple.Render("Date format YYYY-MM-DD\n")
			if !m.TextInput.VirtualCursor() {
				c = m.TextInput.Cursor()
				c.Y += lipgloss.Height(dateFormat + "\n" + headerView)
			}

			strView = strings.Join([]string{
				headerView,
				dateFormat,
				m.TextInput.View(),
				m.styles.StatusMessage.Render(
					fmt.Sprintf("\nThe TravelDate is: %s", m.country.UserData.TravelDate)),
				m.styles.StatusMessage.Render(
					fmt.Sprintf("The TravelLength is: %s", m.country.UserData.TravelEndDate),
				),
				footer,
			}, "\n")
		}

		if !m.isWritingDate {
			strView = strings.Join([]string{
				m.list.View(),
				footer,
			}, "\n")
		}

		if len(m.debugStrs) > 1 {
			strView += strings.Join(m.debugStrs, "\n")
		}

		v := tea.NewView(strView)
		v.AltScreen = true
		v.Cursor = c
		return v
	*/
}

func InitialModel() *Model {
	items := GenerateFromCountryItems()
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)

	ti := textinput.New()
	ti.Placeholder = "Date format YYYY-MM-DD"
	ti.SetValue(time.Now().Format(time.DateOnly))
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 10
	ti.SetWidth(20)

	vp := CreateViewport()

	countryNames, err := utils.GetAllCountriesNames()
	if err != nil {
		log.Fatal(err)
	}

	currentDate := time.Now().Format(time.DateOnly)

	newForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("country").
				Options(huh.NewOptions(countryNames...)...).
				Title("Where are you from?").
				Height(10),

			huh.NewSelect[string]().
				Key("budget").
				Options(huh.NewOptions("Economic", "Moderate", "High Level")...).
				Title("How big is your budget?"),

			huh.NewInput().
				Key("travelDate").
				Title("When are you planning to travel?").
				CharLimit(10).
				Description("Date Format YYYY-MM-DD").
				Placeholder("YYYY-MM-DD").
				Value(&currentDate).
				Validate(func(s string) error {
					rawDateSlice := strings.Split(s, "-")
					if len(rawDateSlice) != 3 {
						return errors.New("Error: Invalid formtat, date must be YYYY-MM-DD")
					}
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
						}
					}
					rawTime, err := time.Parse(time.DateOnly, s)
					if err != nil {
						return fmt.Errorf("Error: Invalid time: %v", err)
					}

					dateNow, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
					if rawTime.Before(dateNow) {
						return errors.New("Invalid date: Traveling date can't be in the past!")
					}

					return nil
				}),

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
				}),

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
				Title("What do you enjoy while traveling?").
				Height(4),
		),
	)

	m := &Model{
		list:          l,
		keys:          createKeyMap(),
		help:          help.New(),
		viewport:      vp,
		TextInput:     ti,
		Form:          newForm,
		progress:      progress.New(progress.WithDefaultBlend()),
		dataMsg:       "Downloading required data...",
		debugStrs:     []string{"\nDEBUG: "},
		questionTitle: "When are you planning to travel?",
	}
	m.list.FilterInput.Placeholder = "testing..."
	m.list.SetFilteringEnabled(true)
	m.UpdateStyles(true)
	m.UpdateDateStyles()
	return m
}
