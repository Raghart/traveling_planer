package bubbletea

import (
	"fmt"

	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
)

func (m *Model) updateRenderer(terminalWidth int) (*glamour.TermRenderer, error) {
	glamourRenderWidth := terminalWidth - m.viewport.Style.GetHorizontalFrameSize() - 3
	styles := styles.DarkStyleConfig

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styles),
		glamour.WithWordWrap(glamourRenderWidth),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to create the renderer: %w", err)
	}
	return renderer, nil
}
