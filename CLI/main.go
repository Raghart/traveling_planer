package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Raghart/traveling_planer/internal/pubsub"
)

func main() {
	fmt.Println("Welcome to Traveling Planner!")
	defer fmt.Println("Good traveling!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CMD > ")
		scanner.Scan()
		switch scanner.Text() {
		case "exit":
			return
		case "who":
			fmt.Println(pubsub.TestingRPC())
		}
	}
}
