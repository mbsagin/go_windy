package main

import (
	"fmt"
	"time"

	"go_windy/utils"
	"go_windy/windy"
)

func main() {
	lat := 41.00
	lon := 28.54
	forecastModel := "iconEu"

	surfaceTemp := windy.GetSurfaceTemperatures(lat, lon, forecastModel)

	for i, ts := range surfaceTemp.Timestamps {
		fmt.Printf("[%s] %.1f Celsius\n", utils.TimestampToUnixTimeUTC(ts).Format(time.RFC822), utils.KelvinToCelsius(surfaceTemp.Temperatures[i]))
	}
}
