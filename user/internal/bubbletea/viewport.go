package bubbletea

import (
	"fmt"

	"charm.land/bubbles/v2/viewport"
	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
	"charm.land/lipgloss/v2"
)

func CreateViewportRenderer() (viewport.Model, *glamour.TermRenderer, error) {
	vp := viewport.New()
	vp.SetWidth(78)
	vp.SetHeight(20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color("62")).
		PaddingRight(2)

	glamourRenderWidth := 78 - vp.Style.GetHorizontalFrameSize() - 3
	styles := styles.DarkStyleConfig

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styles),
		glamour.WithWordWrap(glamourRenderWidth),
	)

	if err != nil {
		return viewport.Model{}, nil, fmt.Errorf("unable to create the renderer: %w", err)
	}

	str, err := renderer.Render("testing!")
	if err != nil {
		return viewport.Model{}, nil, fmt.Errorf("unable to render the text: %w", err)
	}

	vp.SetContent(str)
	return vp, renderer, nil
}
