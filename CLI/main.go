package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/routing"
)

func main() {
	if _, err := tea.NewProgram(routing.InitialModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
