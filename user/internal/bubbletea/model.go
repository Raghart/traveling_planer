package bubbletea

import (
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
		//m.list.SetSize(msg.Width, msg.Height-5)
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

	var (
		country        = "(None)"
		budget         string
		travelDate     string
		travelDuration string
		enviroment     string
		activities     string
	)

	if m.Form.GetString("country") != "" {
		country = fmt.Sprintf("\nFrom: %s\n", m.Form.GetString("country"))
		m.country.UserData.From = country
	}

	if m.Form.GetString("budget") != "" {
		budget = fmt.Sprintf("Budget: %s\n", m.Form.GetString("budget"))
		m.country.UserData.Budget = budget
	}

	if m.Form.GetString("travelDate") != "" {
		travelDate = fmt.Sprintf("Traveling Date: %s\n", m.Form.GetString("travelDate"))

		travelDateTime, _ := time.Parse(time.DateOnly, travelDate)
		m.country.UserData.TravelDate = travelDateTime
	}

	if m.Form.GetString("travelDuration") != "" {
		travelDuration = fmt.Sprintf("Traveling Duration: %s days\n", m.Form.GetString("travelDuration"))

		daysNum, _ := strconv.ParseInt(travelDuration, 10, 32)
		travelEndDate := m.country.UserData.TravelDate.Add(time.Duration(daysNum) * 24 * time.Hour)
		m.country.UserData.TravelEndDate = travelEndDate
	}

	if m.Form.GetBool("enviroment") && m.Form.GetString("travelDuration") != "" {
		enviroment = "Enviroment: Hotter\n"
	} else {
		enviroment = "Enviroment: Colder\n"
	}

	if m.Form.GetString("activities") != "" {
		activities = fmt.Sprintf("Activities: %s\n", m.Form.GetString("activities"))
	}

	const statusWidth = 35
	statusMarginLeft := m.Width - statusWidth - lipgloss.Width(form) - m.styles.Status.GetMarginRight()
	countryStatus := m.styles.Status.
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render("Current Traveling Prefs!"+"\n\n"+
			country+budget+travelDate+travelDuration+enviroment, activities)

	body := lipgloss.JoinHorizontal(lipgloss.Left, form, countryStatus)

	footer := m.AppFrameworkView(m.Form.Help().ShortHelpView(m.Form.KeyBinds()))
	if len(errorSlice) > 0 {
		footer = m.ErrorFrameworkView("")
	}

	//if len(m.debugStrs) > 1 {
	//	strView += strings.Join(m.debugStrs, "\n")
	//}

	return tea.NewView(m.styles.Base.Render(header + "\n" + body + "\n\n" + footer))
}

func InitialModel() *Model {
	m := &Model{
		styles:   NewStyles(true),
		keys:     createKeyMap(),
		help:     help.New(),
		progress: progress.New(progress.WithDefaultBlend()),
		dataMsg:  "Downloading required data...",
	}

	err := m.CreateNewForm()
	if err != nil {
		log.Print("Something went wrong!")
		log.Fatal(err)
	}

	return m
}
