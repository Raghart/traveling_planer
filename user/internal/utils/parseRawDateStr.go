package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRawDateStr(rawDateStr string) (string, error) {
	rawTravelSlice := strings.Split(rawDateStr, "-")
	if len(rawTravelSlice) != 3 {
		return "", fmt.Errorf("Invalid date format: %v", rawDateStr)
	}

	parsedTravelSlice := []string{}
	for _, dateStr := range rawTravelSlice {
		dateNum, err := strconv.ParseInt(dateStr, 10, 32)
		if err != nil {
			return "", fmt.Errorf("unable to parse the value: %v", dateStr)
		}
		parsedTravelSlice = append(parsedTravelSlice, fmt.Sprintf("%02d", dateNum))
	}

	return strings.Join(parsedTravelSlice, "-"), nil
}
