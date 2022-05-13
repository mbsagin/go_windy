package utils

import (
	"log"
	"os"
	"time"
)

// Parse timestamp to time.Unix UTC format
func TimestampToUnixTimeUTC(timestamp int64) time.Time {
	return time.Unix(timestamp / 1000, 0)
}

// ℃ = K - 273.15
var kelvinConstant float64 = 273.15

// Convert kelvin to celsius
func KelvinToCelsius (kelvin float64) float64 {
	return kelvin - kelvinConstant
}

// Gets environment variable with the key param
func GetEnvVariable(key string) string {
	envVariable, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Env key could not found: %s", key)
		return key
	}

	return envVariable
}