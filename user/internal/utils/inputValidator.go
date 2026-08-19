package utils

import (
	"fmt"
	"time"
)

func IsValidInput(input string) error {
	parsedDate, err := time.Parse(time.DateOnly, input)
	if err != nil {
		return fmt.Errorf("this string '%s' doesn't match the expected format", input)
	}
	if !parsedDate.After(time.Now()) {
		return fmt.Errorf("Error: Traveling date can't be in the past!")
	}
	return nil
}
