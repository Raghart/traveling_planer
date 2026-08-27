package bubbletea

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"charm.land/huh/v2"
)

type state int

const (
	ProcessingForm state = iota
	FormFilled
)

func (m *Model) CreateNewForm() error {
	countryNames := getCountryNames()

	currentDate := time.Now().Format(time.DateOnly)
	defaultDaysNum := "7"

	newForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("country").
				Options(countryNames...).
				Title("Where are you from?").
				Inline(true).
				Description(m.styles.Help.Render("Select your origin's country")),

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
						return m.styles.Help.Render("Select your budget")
					}
				}, &m.country.UserData.Budget),

			huh.NewInput().
				Key("travelDate").
				Title("When are you planning to travel?").
				CharLimit(10).
				Description(m.styles.Help.Render("Date Format YYYY-MM-DD | YYYY-M-DD")).
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
				Value(&defaultDaysNum).
				Validate(func(s string) error {
					_, err := strconv.ParseInt(s, 10, 32)
					if err != nil {
						return fmt.Errorf("Error: number is not a valid int: %v", err)
					}

					return nil
				}).
				Description(m.styles.Help.Render("Select the number of days you're planning to travel")),

			huh.NewConfirm().
				Key("enviroment").
				Title("Do you prefer hotter or colder enviroments?").
				Affirmative("Hotter").
				Negative("Colder").
				Value(&m.country.UserData.PrefersWarmth),

			huh.NewMultiSelect[string]().
				Key("activities").
				Options(
					huh.NewOption("Nature-wilderness", "nature"),
					huh.NewOption("City-life style", "city"),
					huh.NewOption("History-culture", "history"),
				).
				Title("What do you enjoy doing while traveling?").
				Description(m.styles.Help.Render("Select the activities you enjoy doing!")).
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
