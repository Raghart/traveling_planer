package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseRawDateStr(rawDateStr string) (string, error) {
	rawTravelSlice := strings.Split(rawDateStr, "-")
	if len(rawTravelSlice) != 3 {
		return "", fmt.Errorf("Invalid date format: %v", rawDateStr)
	}

	parsedTravelSlice := []string{}
	for idx, dateStr := range rawTravelSlice {
		dateNum, err := strconv.ParseInt(dateStr, 10, 32)
		if err != nil {
			return "", fmt.Errorf("unable to parse the value: %v", dateStr)
		}
		switch idx {
		case 0:
			if dateNum < int64(time.Now().Year()) {
				return "", fmt.Errorf("Invalid year, traveling date can't be in the past!")
			}
			if dateNum > 3000 {
				return "", fmt.Errorf("Invalid year, traveling date is too far into the future!")
			}
		}
		parsedTravelSlice = append(parsedTravelSlice, fmt.Sprintf("%02d", dateNum))
	}

	return strings.Join(parsedTravelSlice, "-"), nil
}
