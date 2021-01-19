package helper

import "os"

const envKey = "ARTENV"

var currentEnv = os.Getenv(envKey)

func GetEnv() string {
	if currentEnv == "" {
		currentEnv = "development"
		return "development"
	}
	return currentEnv
}

func IsDevelopment() bool {
	if currentEnv == "development" {
		return true
	}
	return false
}
