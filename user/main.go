package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Raghart/traveling_planer/internal/bubbletea"
	"github.com/Raghart/traveling_planer/internal/types"
)

func main() {
	res, err := http.Get("https://countries.dev/countries")
	if err != nil {
		log.Fatalf("error while trying to retrieve the country data: %v", err)
	}

	jsonData := &types.CountriesJSON{}
	bodyData, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("error while trying to read the body: %v", err)
	}

	err = json.Unmarshal(bodyData, jsonData)
	if err != nil {
		log.Fatalf("unable to unmarshal the body: %v", err)
	}
	fmt.Printf("There are currently %d countries in the json!\n", len(*jsonData))

	return
	if _, err := tea.NewProgram(bubbletea.InitialModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
