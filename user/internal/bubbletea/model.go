package bubbletea

import (
	"fmt"
	"slices"
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
	"charm.land/lipgloss/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
)

type Model struct {
	keys          KeyMap
	list          list.Model
	styles        Styles
	help          help.Model
	country       routing.CountryManager
	progress      progress.Model
	viewport      viewport.Model
	renderer      *glamour.TermRenderer
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
	return tea.Batch(
		tea.RequestBackgroundColor,
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				userStrTravelDate := m.TextInput.Value()
				userTravelDate, err := time.Parse(time.DateOnly, userStrTravelDate)

				if err != nil {
					m.debugStrs = append(m.debugStrs, fmt.Sprintf("ERROR: %v", err))
					return m, nil
				}

				if userTravelDate.Before(time.Now()) {
					m.debugStrs = append(m.debugStrs, "Error: Travel Date can't be in the past")
					return m, nil
				}

				m.country.UserData.TravelDate = userTravelDate
				return m, nil
			}

			if ok && m.country.From != "" && m.country.Destination != "" {

				return m, tea.Quit
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
}

func (m *Model) View() tea.View {
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
		headerView := m.styles.Title.Render("When are you planning to travel?")
		if !m.TextInput.VirtualCursor() {
			c = m.TextInput.Cursor()
			c.Y += lipgloss.Height(headerView)
		}

		strView = strings.Join([]string{
			headerView,
			m.TextInput.View(),
			m.styles.StatusMessage.Render(
				fmt.Sprintf("\nThe UserDate is: %s", m.country.UserData.TravelDate)),
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
}

func InitialModel() *Model {
	items := GenerateFromCountryItems()
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Where are you from?"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)

	ti := textinput.New()
	ti.Placeholder = "Date format YYYY-MM-DD"
	currentTime := time.Now()
	ti.SetValue(fmt.Sprintf("%d-%02d-%02d", currentTime.Year(), currentTime.Month(), currentTime.Day()))
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 10
	ti.SetWidth(20)

	vp := CreateViewport()

	m := &Model{
		list:      l,
		keys:      createKeyMap(),
		help:      help.New(),
		viewport:  vp,
		TextInput: ti,
		progress:  progress.New(progress.WithDefaultBlend()),
		dataMsg:   "Downloading required data...",
		debugStrs: []string{"\nDEBUG: "},
	}
	m.list.FilterInput.Placeholder = "testing..."
	m.list.SetFilteringEnabled(true)
	m.UpdateStyles(true)
	m.UpdateDateStyles()
	return m
}
