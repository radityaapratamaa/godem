package helper

import (
	"github.com/kodekoding/phastos/go/env"
	"os"
)

var currentEnv = os.Getenv(env.Name)

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
