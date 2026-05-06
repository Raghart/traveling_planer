package utils

import "log"

func FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s : %v", message, err)
	}
}
