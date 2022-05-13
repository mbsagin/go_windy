package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"go_windy/request"
	"go_windy/utils"
)

var apiKey string
var apiUrl string

type requestModel struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Model string `json:"model"`
	Parameters []string `json:"parameters"`
	Levels []string `json:"levels"`
	Key string `json:"key"`
}

type surfaceTemperatureModel struct {
	Timestamps []int64 `json:"ts"`
	Temperatures []float64 `json:"temp-surface"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Env could not found: %s", err)
	}

	apiKey = utils.GetEnvVariable("API_KEY")
	apiUrl = utils.GetEnvVariable("API_URL")
}

func main() {
	lat := 41.00
	lon := 28.54
	forecastModel := "iconEu"

	surfaceTemps := getSurfaceTemperatures(lat, lon, forecastModel)

	for i, ts := range surfaceTemps.Timestamps {
		fmt.Printf("[%s] %.1f Celcius\n", utils.TimestampToUnixTimeUTC(ts).Format(time.RFC822), utils.KelvinToCelsius(surfaceTemps.Temperatures[i]))
	}
}

func getSurfaceTemperatures(lat float64, lon float64, forecastModel string) surfaceTemperatureModel {
	tempModel := surfaceTemperatureModel{}

	model := requestModel {
		lat, 
		lon, 
		forecastModel, 
		[]string {"temp"}, 
		[]string {"surface"}, 
		apiKey }

	modelBytes, err := json.Marshal(model)

	if err != nil {
		log.Fatal(err)
	}

	responseBody := request.MakePostRequest(apiUrl, modelBytes)

	if err := json.Unmarshal(responseBody, &tempModel); err != nil {
		log.Fatal(err)
	}

	return tempModel
}