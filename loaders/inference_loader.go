package loaders

import (
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/csv"
	"encoding/json"
)

type Scalers struct {
	MinValue map[string]float32 `json:"min_value"`
	MaxValue map[string]float32 `json:"max_value"`
}

func ScalersLoader(filename string) (Scalers, error) {
	var scalers_data Scalers

	scalers_json, err := ioutil.ReadFile(filename)
	if err != nil { return scalers_data, err }

	err = json.Unmarshal(scalers_json, &scalers_data)
	if err != nil { return scalers_data, err }

	return scalers_data, nil
}

func DatasetLoader(filepath string) ([][]interface{}, error) {
	csv_file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)
	csv_records, err := csv_reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csv_data [][]interface{}
	for _, csv_row := range csv_records[1:] {
		var row_data []interface{}

		// kolom pertama tetap string (Date)
		row_data = append(row_data, csv_row[0])

		// sisanya float32
		for _, csv_row_val := range csv_row[1:] {
			val, err := strconv.ParseFloat(csv_row_val, 32)
			if err != nil {
				return nil, fmt.Errorf("error parsing value '%s': %v", csv_row_val, err)
			}
			row_data = append(row_data, float32(val))
		}

		csv_data = append(csv_data, row_data)
	}

	return csv_data, nil
}


