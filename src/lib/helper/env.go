package helper

import (
	"os"

	"github.com/kodekoding/phastos/v2/go/env"
)

var currentEnv = os.Getenv(env.Name)

func GetEnv() string {
	if currentEnv == "" {
		currentEnv = env.LocalEnv
	}
	return currentEnv
}

func IsLocal() bool {
	return currentEnv == env.LocalEnv
}
