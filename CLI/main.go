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
		case "curr":
			if len(inputSlice) != 3 {
				fmt.Println("Invalid curr cmd usage. Example: CMD > curr USD CAD")
				continue
			}

			fromCurr := inputSlice[1]
			toCurr := inputSlice[2]
			result := pubsub.AskCurrency(fromCurr, toCurr)
			fmt.Printf("Currency value from '%s' to '%s': %f\n", fromCurr, toCurr, result)
		}
	}
}
