package loaders

import (
	"os"
	"encoding/json"
)

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