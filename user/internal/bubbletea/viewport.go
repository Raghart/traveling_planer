package bubbletea

import (
	"charm.land/bubbles/v2/viewport"
	"charm.land/lipgloss/v2"
)

func CreateViewport() viewport.Model {
	vp := viewport.New()
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color("62"))

	return vp
}
