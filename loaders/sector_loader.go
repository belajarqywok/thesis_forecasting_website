package loaders

import (
	"os"
	"encoding/json"
)

func SectorLoader(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data SectorsWrapper
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Sectors, nil
}
