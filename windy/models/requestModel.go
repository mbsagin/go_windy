package windy

type RequestModel struct {
	Lat        float64  `json:"lat"`
	Lon        float64  `json:"lon"`
	Model      string   `json:"model"`
	Parameters []string `json:"parameters"`
	Levels     []string `json:"levels"`
	Key        string   `json:"key"`
}