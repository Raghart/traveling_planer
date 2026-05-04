package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Traveling Planner!")
	defer fmt.Println("Good traveling!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CMD > ")
		scanner.Scan()
		if scanner.Text() == "exit" {
			break
		}
	}
}
