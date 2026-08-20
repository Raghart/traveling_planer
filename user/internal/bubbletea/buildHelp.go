package bubbletea

import "strings"

func (m *Model) buildDynamicHelp() string {
	if m.isWritingDate {
		return strings.Join([]string{
			"← / →: Navigate",
			"Backspace: Delete",
			"Enter: Select the date",
			"Q/Esc: Exit",
		}, " • ")
	}

	return strings.Join([]string{
		"↑/↓/←/→: Navigate",
		"Enter: Select",
		"Filter: /",
		"Q/Esc: Exit",
	}, " • ")
}
