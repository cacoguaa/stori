package utils

import (
	"log"
	"os"
	"strconv"
)

func Get(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Fatal("Missing environment variable: " + envName + " must be set")
	}

	return value
}

func GetInt(envName string) int {
	value := Get(envName)
	parsedVal, err := strconv.Atoi(value)

	if err != nil {
		log.Fatal("Can't parse environment variable " + envName + " to int")
	}

	return parsedVal
}
