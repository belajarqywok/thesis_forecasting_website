package loaders

import (
	"os"
	"encoding/json"
)

type Indicator struct {
	Date            string  `json:"date"`
	FullDate        string  `json:"full_date"`
	MFI            float64  `json:"MFI"`
	RSI            float64  `json:"RSI"`
	MACD           float64  `json:"MACD"`
}

type IndicatorsWrapper struct {
	Indicators   []Indicator   `json:"indicators"`
}

func IndicatorLoader(path string) ([]Indicator, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data IndicatorsWrapper
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Indicators, nil
}