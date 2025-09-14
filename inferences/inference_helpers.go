package inferences

import (
	"sync"
	loaders "thesis_forecasting_website/loaders"
)

type StockPrice struct {
	Date  string  `json:"date"`
	Price float32 `json:"price"`
}

func Denormalization(data, minValue, maxValue float32) float32 {
	return (data * (maxValue - minValue)) + minValue
}

func InferenceLoader(inferenceDataPath, scalersDataPath string) (
	[][]interface{}, loaders.Scalers, []error,
) {
	var (
    inferenceData   [][]interface{}
    scalersData     loaders.Scalers
  )

	errChannel := make(chan error, 2)

	var wgDatasetScalerLoader sync.WaitGroup
  wgDatasetScalerLoader.Add(2)

	go func() {
    defer wgDatasetScalerLoader.Done()
    tempData, err := loaders.DatasetLoader(inferenceDataPath)
    if err != nil {
			errChannel <- err
			return
		}
		inferenceData = tempData
  }()

	go func() {
		defer wgDatasetScalerLoader.Done()
    tempData, err := loaders.ScalersLoader(scalersDataPath)
    if err != nil {
			errChannel <- err
			return
		}
		scalersData = tempData
	}()

	wgDatasetScalerLoader.Wait()
	close(errChannel)

	var errors []error
  for err := range errChannel {
    errors = append(errors, err)
  }

  if len(errors) > 0 {
    return nil, loaders.Scalers{}, errors
  }

  return inferenceData, scalersData, nil
}