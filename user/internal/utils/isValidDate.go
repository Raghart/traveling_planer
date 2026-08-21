package utils

import (
	"time"
)

func IsInvalidDate(date time.Time) bool {
	actualDate, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	if date.Before(actualDate) {
		return true
	}

	if date.Format(time.DateOnly) == time.Now().Format(time.DateOnly) {
		return false
	}

	return false
}
