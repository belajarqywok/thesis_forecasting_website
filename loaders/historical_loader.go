package loaders

import (
	"os"
	"encoding/json"
)

type Historical struct {
	Date           string   `json:"date"`
	FullDate       string   `json:"full_date"`
	Open          float64   `json:"open"`
	High          float64   `json:"high"`
	Low           float64   `json:"low"`
	Close         float64   `json:"close"`
  Volume        float64   `json:"volume"`
}

type HistoricalsWrapper struct {
	Historicals   []Historical   `json:"historicals"`
}

func HistoricalLoader(path string) ([]Historical, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data HistoricalsWrapper
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Historicals, nil
}
