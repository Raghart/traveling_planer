package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Raghart/traveling_planer/internal/types"
)

func GetAllCountriesData() (types.CountriesJSON, error) {
	res, err := http.Get("https://countries.dev/countries")
	if err != nil {
		return nil, fmt.Errorf("error while trying to retrieve the country data: %v", err)
	}

	jsonData := &types.CountriesJSON{}
	bodyData, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error while trying to read the body: %v", err)
	}

	err = json.Unmarshal(bodyData, jsonData)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal the body: %v", err)
	}
	return *jsonData, nil
}
