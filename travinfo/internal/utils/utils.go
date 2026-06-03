package utils

import (
	"encoding/json"
	"log"
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

func FailsOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
