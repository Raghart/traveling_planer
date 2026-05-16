package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Raghart/traveling_planer/internal/pubsub"
	"github.com/Raghart/traveling_planer/internal/utils"
)

func main() {
	fmt.Println("Welcome to Traveling Planner!")
	defer fmt.Println("Good traveling!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CMD > ")
		scanner.Scan()
		inputSlice := strings.Split(scanner.Text(), " ")

		switch inputSlice[0] {
		case "exit":
			return
		case "test":
			fmt.Println(pubsub.TestingRPC())
		case "temp":
			if len(inputSlice) != 2 {
				fmt.Println("Invalid temp cmd usage. Example: CMD > temp Canada")
			}
			country := inputSlice[1]
			dailyTemps := pubsub.AskTemperature(country)
			utils.PresentWeather(dailyTemps)

		case "curr":
			if len(inputSlice) != 3 {
				fmt.Println("Invalid curr cmd usage. Example: CMD > curr USD CAD")
				continue
			}

			fromCurr := inputSlice[1]
			toCurr := inputSlice[2]
			result := pubsub.AskCurrency(fromCurr, toCurr)
			fmt.Printf("Currency value from %s to '%s': %.2f\n", fromCurr, toCurr, result)
		}
	}
}
