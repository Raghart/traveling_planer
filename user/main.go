package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(bubbletea.InitialModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
