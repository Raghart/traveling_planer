package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Raghart/traveling_planer/internal/pubsub"
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
			pubsub.AskTemperature(country)

		case "curr":
			if len(inputSlice) != 2 {
				fmt.Println("Invalid curr cmd usage. Example: CMD > curr CAD")
				continue
			}

			toCurr := inputSlice[1]
			result := pubsub.AskCurrency(toCurr)
			fmt.Printf("Currency value from USD to '%s': %.2f\n", toCurr, result)
		}
	}
}
