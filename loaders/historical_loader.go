package loaders

import (
	"os"
	"encoding/json"
)

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
