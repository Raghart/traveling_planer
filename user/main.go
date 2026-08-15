package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/bubbletea"
	"github.com/Raghart/traveling_planer/internal/utils"
)

func main() {
	countryList, err := utils.GetAllCountriesData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are currently %d countries in the json!\n", len(*countryList))
	return
	if _, err := tea.NewProgram(bubbletea.InitialModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
