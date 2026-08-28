package bubbletea

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

type state int

const (
	ProcessingForm state = iota
	FormFilled
)

func (m *Model) CreateNewForm() error {
	countryNames := getCountryNames()

	newForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("country").
				Options(countryNames...).
				Title("Where are you from?").
				Inline(true).
				Description(m.styles.Help.Render("Select your origin's country")).
				Value(&m.country.UserData.From),

			huh.NewSelect[string]().
				Key("budget").
				Options(huh.NewOptions("Economic", "Moderate", "High Level")...).
				Title("How big is your budget?").
				Value(&m.country.UserData.Budget).
				DescriptionFunc(func() string {
					switch m.country.UserData.Budget {
					case "Economic":
						return m.styles.Help.Render("Budget equal or lower than 5.000$")
					case "Moderate":
						return m.styles.Help.Render("Budget Between 5.000$ and 10.000$")
					case "High Level":
						return m.styles.Help.Render("Budget highter than 10.000$")
					default:
						return m.styles.Help.Render("Budget saved!")
					}
				}, &m.country.UserData.Budget),

			huh.NewInput().
				Key("travelDate").
				Title("When are you planning to travel?").
				CharLimit(10).
				Description(m.styles.Help.Render("Date Format YYYY-MM-DD | YYYY-M-DD")).
				Placeholder("YYYY-MM-DD").
				Value(&m.country.UserData.TravelDate).
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
				Value(&m.country.UserData.TravelDays).
				Validate(func(s string) error {
					parsedInt, err := strconv.ParseInt(s, 10, 32)
					if err != nil {
						return fmt.Errorf("Error: number is not a valid int: %v", err)
					}

					if parsedInt <= 0 {
						return errors.New("Error: It's imposible to travel in 0 or negative days!")
					}

					return nil
				}).
				Description(m.styles.Help.Render("Select the number of days you're planning to travel")),

			huh.NewSelect[string]().
				Key("enviroment").
				Title("Do you prefer hotter or colder enviroments?").
				Options(huh.NewOptions("Hotter", "Colder")...).
				Value(&m.country.UserData.PrefersWarmth),

			huh.NewMultiSelect[string]().
				Key("activities").
				Options(
					huh.NewOption("Nature-wilderness", "Nature-wilderness"),
					huh.NewOption("City-life style", "City-life style"),
					huh.NewOption("History-culture", "History-culture"),
				).
				Title("Which areas would you prefer to explore?").
				Height(5).
				Value(&m.country.UserData.PreferedActivities),
		),
	).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeFunc(huh.ThemeDracula))

	m.Form = newForm
	return nil
}

func (m *Model) GenerateStatusView(form string) string {
	var (
		countryStat = m.styles.StatusHeader.Render("Waiting for the data!")
		country,
		budget, budgetStat,
		travelDate, travelDateStat,
		travelDuration, travelDurStat,
		enviroment, enviromentStat,
		activities string
	)

	if m.Form.GetString("country") != "" || country != m.Form.GetString("country") {
		country = m.Form.GetString("country")
		countryStat = m.styles.StatusHeader.Render("From: ") +
			fmt.Sprintf("%s\n", m.Form.GetString("country"))
	}

	if m.Form.GetString("budget") != "" || budget != m.Form.GetString("budget") {
		budget = m.Form.GetString("budget")
		budgetStat = m.styles.StatusHeader.Render("Budget: ") +
			fmt.Sprintf("%s\n", m.Form.GetString("budget"))
	}

	if m.Form.GetString("travelDate") != "" || travelDate != m.Form.GetString("travelDate") {
		travelDate = m.Form.GetString("travelDate")
		travelDateStat = m.styles.StatusHeader.Render("Traveling Date: ") + fmt.
			Sprintf("%s\n", m.Form.GetString("travelDate"))
	}

	if m.Form.GetString("travelDuration") != "" || travelDuration != m.Form.GetString("travelDuration") {
		travelDuration = m.Form.GetString("travelDuration")
		travelDurStat = m.styles.StatusHeader.Render("Traveling Duration: ") + fmt.Sprintf(
			"%s days\n", m.Form.GetString("travelDuration"))
	}

	if m.Form.GetString("enviroment") != "" || enviroment != m.Form.GetString("enviroment") {
		enviroment = m.Form.GetString("enviroment")
		enviromentStat = m.styles.StatusHeader.Render("Enviroment: ") + fmt.
			Sprintf("%s\n", m.Form.GetString("enviroment"))
	}

	if len(m.country.UserData.PreferedActivities) > 0 {
		activities = m.styles.StatusHeader.Render("Activities:") + "\n"
		for _, act := range m.country.UserData.PreferedActivities {
			activities += m.styles.StatusText.Render("- ") + fmt.Sprintf("%s\n", act)
		}
	}

	const statusWidth = 35
	dataHeader := m.styles.HeaderText.Render("Traveling Data Gathered")

	return m.styles.Status.
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(m.Width - statusWidth - lipgloss.Width(form) - m.styles.Status.GetMarginRight()).
		Render(
			dataHeader + "\n\n" +
				countryStat +
				budgetStat +
				travelDateStat +
				travelDurStat +
				enviromentStat +
				activities,
		)
}
