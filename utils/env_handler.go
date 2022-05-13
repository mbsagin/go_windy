package utils

import (
	"log"
	"os"
)

// Gets environment variable with the key param
func GetEnvVariable(key string) string {
	envVariable, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Env key could not found: %s", key)
		return key
	}

	return envVariable
}