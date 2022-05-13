package utils

import "time"

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