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

	surfaceTemps := windy.GetSurfaceTemperatures(lat, lon, forecastModel)

	for i, ts := range surfaceTemps.Timestamps {
		fmt.Printf("[%s] %.1f Celcius\n", utils.TimestampToUnixTimeUTC(ts).Format(time.RFC822), utils.KelvinToCelsius(surfaceTemps.Temperatures[i]))
	}
}