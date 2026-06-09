package bubbletea

import (
	"charm.land/bubbles/v2/viewport"
	"charm.land/lipgloss/v2"
)

func CreateViewport() viewport.Model {
	vp := viewport.New()

	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 1).
		MarginLeft(2).
		MarginRight(2)

	return vp
}
