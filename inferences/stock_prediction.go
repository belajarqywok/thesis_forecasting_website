package inferences

import (
	"fmt"
	"log"
	"time"
	"sync"

	onnxruntime "github.com/belajarqywok/onnxruntime_go"
)


/*

  -- Stock Prediction -- 

  Writer : Al-Fariqy Raihan Azhwar
  NPM    : 202143501514
  Class  : R8Q
  Email  : alfariqyraihan@gmail.com


*/


var once sync.Once
var isInit bool

func StockPrediction(issuer string, days int) (
	[]StockPrice, []StockPrice, error,
) {
	once.Do(func() {
		runtimePath := "./onnxruntime-linux-x64-1.21.0/lib/libonnxruntime.so"
		onnxruntime.SetSharedLibraryPath(runtimePath)

		if err := onnxruntime.InitializeEnvironment(); err != nil {
			log.Fatal("Error initializing ONNX runtime: ", err)
		}

		isInit = true
	})
	if !isInit { log.Fatal("ONNX runtime not initialized") }

	inferenceDataPath := fmt.Sprintf("./indonesia_stocks/modeling_datas/%s.csv", issuer)
	scalersDataPath   := fmt.Sprintf("./indonesia_stocks/min_max/%s.json", issuer)
	data, scalers, errors := InferenceLoader(inferenceDataPath, scalersDataPath)
	if len(errors) > 0 {
		for _, e := range errors { log.Println("Error:", e) }
		return []StockPrice{}, []StockPrice{}, fmt.Errorf(
			"multiple errors occurred: %v", errors,
		)
	}

	nData := 30
	if len(data) < nData { nData = len(data) }
	lastActualData := data[len(data) - nData:]

	var actuals []StockPrice
	for _, row := range lastActualData {
		date := row[0].(string)
		closeVal := row[1].(float32)

		closePrice := Denormalization(
			closeVal,
			scalers.MinValue["Close"],
			scalers.MaxValue["Close"],
		)

		actuals = append(actuals, StockPrice{
			Date:  date,
			Price: float32(closePrice),
		})
	}

	sequenceLength   := int64(60)
	featureSize      := int64(5)
	lastSequenceData := data[len(data) - int(sequenceLength):]

	inputData  := make([]float32, sequenceLength * featureSize)
	for idxRow, row := range lastSequenceData {
		for idxFeature := 1; idxFeature <= int(featureSize); idxFeature++ {
			valueFeature, ok := row[idxFeature].(float32)
			if !ok {log.Fatalf(
				"expected float32 at row %d col %d, got %T", 
				idxRow, idxFeature, row[idxFeature],
			)}

			inputData[idxRow*int(featureSize) + (idxFeature-1)] = valueFeature
		}
	}

	inputShape := onnxruntime.NewShape(1, sequenceLength, featureSize)
	inputTensor, err := onnxruntime.NewTensor(inputShape, inputData)
	if err != nil { log.Fatalf("error creating input tensor: %v", err) }

	outputShape := onnxruntime.NewShape(1, 1)
	outputTensor, err := onnxruntime.NewEmptyTensor[float32](outputShape)
	if err != nil { log.Fatalf("error creating output tensor: %v", err) }

	session, err := onnxruntime.NewAdvancedSession(
		fmt.Sprintf("./models/%s.onnx", issuer),
		[]string{"input"},   []string{"output"},
		[]onnxruntime.ArbitraryTensor{inputTensor},
		[]onnxruntime.ArbitraryTensor{outputTensor}, nil,
	)
	if err != nil { log.Fatalf("error initializing ONNX session: %v", err) }

	var predicted []StockPrice
	lastDate, _ := time.Parse("2006-01-02", actuals[len(actuals)-1].Date)
	for i := 0; i < days; i++ {
		if err := session.Run(); err != nil {
			log.Fatalf("Error running model: %v", err)
		}

		predictedClose := outputTensor.GetData()[0]
		denormPrice := Denormalization(
			predictedClose, 
			scalers.MinValue["Close"], 
			scalers.MaxValue["Close"],
		)

		lastDate  = lastDate.AddDate(0, 0, 1)
		predicted = append(predicted, StockPrice{
			Date:  lastDate.Format("2006-01-02"),
			Price: float32(denormPrice),
		})

		copy(inputData, inputData[int(featureSize):])
		inputData[len(inputData)-1] = predictedClose
	}

	return actuals, predicted, nil
}
