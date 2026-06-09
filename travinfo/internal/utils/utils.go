package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/Raghart/traveling_planer/travinfo/internal/types"
	"github.com/charmbracelet/x/term"
)

func StructToDict(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	err = json.Unmarshal(data, &res)
	return res, nil
}

func GetTerminalWidth() int {
	width, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return 80
	}
	return width
}

func FailsOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func GetPhotoUrls(countryName string) ([]string, error) {
	baseUrl := "https://api.unsplash.com/search/photos"
	params := url.Values{}
	params.Add("query", countryName)
	params.Add("order_by", "relevant")
	params.Add("client_id", os.Getenv("ACESSKEY"))

	photoUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	res, err := http.Get(photoUrl)
	if err != nil || res.StatusCode > 400 {
		return nil, fmt.Errorf("error trying to get the photos: %w", err)
	}

	defer res.Body.Close()

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read the body of the photos: %w", err)
	}

	jsonPhotos := &types.PhotoJsonURL{}
	err = json.Unmarshal(bodyData, jsonPhotos)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal the json: %w", err)
	}

	photosUrls := []string{}

	for _, data := range jsonPhotos.Results {
		photosUrls = append(photosUrls, data.Urls.Raw)
	}

	return photosUrls, nil
}
