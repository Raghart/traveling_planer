package utils

import (
	"encoding/json"
	"log"
	"os"

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
