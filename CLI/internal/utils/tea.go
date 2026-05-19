package utils

import (
	"charm.land/bubbles/v2/list"
	"github.com/Raghart/traveling_planer/internal/routing"
)

func InitialModel() routing.Model {
	items := []list.Item{
		routing.Item("Argentina"),
		routing.Item("Bolivia"),
		routing.Item("Brazil"),
		routing.Item("Canada"),
		routing.Item("Chile"),
		routing.Item("Colombia"),
		routing.Item("Costa Rica"),
		routing.Item("Cuba"),
		routing.Item("Dominica"),
		routing.Item("Dominican Republic"),
		routing.Item("Grenada"),
		routing.Item("French Guiana"),
		routing.Item("Guyana"),
		routing.Item("Saint Lucia"),
		routing.Item("Honduras"),
		routing.Item("Mexico"),
		routing.Item("Nicaragua"),
		routing.Item("Panama"),
		routing.Item("Peru"),
		routing.Item("Puerto Rico"),
		routing.Item("Paraguay"),
		routing.Item("Suriname"),
		routing.Item("El Salvador"),
		routing.Item("Trinidad and Tobago"),
		routing.Item("United States"),
		routing.Item("Uruguay"),
		routing.Item("Venezuela"),
		routing.Item("Guatemala"),
		routing.Item("Belize"),
		routing.Item("Jamaica"),
		routing.Item("Haiti"),
		routing.Item("Bahamas"),
		routing.Item("Barbados"),
		routing.Item("Saint Kitts and Nevis"),
		routing.Item("Antigua and Barbuda"),
	}

	const defaultWidth = 20

	l := list.New(items, routing.ItemDelegate{}, defaultWidth, 14)
	l.Title = "Where are you From?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	m := routing.Model{
		List: l,
	}
	m.UpdateStyles(true)
	return m
}
