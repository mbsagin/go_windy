package windy

import (
	"encoding/json"
	"log"
	
	"github.com/joho/godotenv"
	
	"go_windy/request"
	"go_windy/utils"
	"go_windy/windy/models"
)

var apiKey string
var apiUrl string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Env could not found: %s", err)
	}

	apiKey = utils.GetEnvVariable("API_KEY")
	apiUrl = utils.GetEnvVariable("API_URL")
}

func GetSurfaceTemperatures(lat float64, lon float64, forecastModel string) windy.SurfaceTemperatureModel {
	tempModel := windy.SurfaceTemperatureModel{}

	model := windy.RequestModel{
		Lat: lat,
		Lon: lon,
		Model: forecastModel,
		Parameters: []string{"temp"},
		Levels: []string{"surface"},
		Key: apiKey}

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