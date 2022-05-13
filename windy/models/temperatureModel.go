package windy

type SurfaceTemperatureModel struct {
	Timestamps   []int64   `json:"ts"`
	Temperatures []float64 `json:"temp-surface"`
}