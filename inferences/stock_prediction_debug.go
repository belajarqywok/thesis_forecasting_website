package inferences

import (
	"fmt"
	"log"
	"time"
	"sync"

	helpers     "thesis_forecasting_website/helpers"
	onnxruntime "github.com/belajarqywok/onnxruntime_go"
)


/*

  -- Stock Prediction [ Debug ] -- 

  Writer : Al-Fariqy Raihan Azhwar
  NPM    : 202143501514
  Class  : R8Q
  Email  : alfariqyraihan@gmail.com


*/


var onceDebug sync.Once
var isInitDebug bool

func StockPredictionDebug(issuer string, days int) ([]StockPrice, []StockPrice, error) {
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeLoadOnnxRuntime := time.Now()
	onceDebug.Do(func() {
		runtimePath := "./onnxruntime-linux-x64-1.21.0/lib/libonnxruntime.so"
		onnxruntime.SetSharedLibraryPath(runtimePath)

		if err := onnxruntime.InitializeEnvironment(); err != nil {
			log.Fatal("Error initializing ONNX runtime: ", err)
		}

		isInitDebug = true
	})
	if !isInitDebug { log.Fatal("ONNX runtime not initialized") }
	
	elapsedTimeGetLoadOnnxRuntime:= time.Since(startTimeLoadOnnxRuntime)
	fmt.Printf(
		"[ Time ] Load ONNX Runtime : %.9f s\n\n", 
		elapsedTimeGetLoadOnnxRuntime.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetInferenceScalerData := time.Now()

	inferenceDataPath := fmt.Sprintf("./indonesia_stocks/modeling_datas/%s.csv", issuer)
	scalersDataPath   := fmt.Sprintf("./indonesia_stocks/min_max/%s.json", issuer)
	data, scalers, errors := InferenceLoader(inferenceDataPath, scalersDataPath)
	if len(errors) > 0 {
		for _, e := range errors { log.Println("Error:", e) }
		return []StockPrice{}, []StockPrice{}, fmt.Errorf("multiple errors occurred: %v", errors)
	}

	elapsedTimeGetInferenceScalerData := time.Since(startTimeGetInferenceScalerData)
	fmt.Printf(
		"[ Time ] Get Inference Scaler Data : %.9f s\n\n", 
		elapsedTimeGetInferenceScalerData.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(data) -- ]")
	fmt.Println(data)
	fmt.Println()

	fmt.Println("[ -- fmt.Println(len(data)) -- ]")
	fmt.Println(len(data))
	fmt.Println()

	fmt.Println("[ -- fmt.Println(scalers) -- ]")
	fmt.Println(scalers)
	fmt.Println()

	fmt.Println("[ -- fmt.Println(errors) -- ]")
	fmt.Println(errors)
	fmt.Println()

	nData := 30
	if len(data) < nData { nData = len(data) }
	lastActualData := data[len(data) - nData:]


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetActualsData := time.Now()

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

	elapsedTimeGetActualsData := time.Since(startTimeGetActualsData)
	fmt.Printf(
		"[ Time ] Get Actuals Data : %.9f s\n\n", 
		elapsedTimeGetActualsData.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(actuals) -- ]")
	fmt.Println(actuals)
	fmt.Println()

	fmt.Println("[ -- fmt.Println(len(actuals)) -- ]")
	fmt.Println(len(actuals))
	fmt.Println()

	// prepare input for model
	sequenceLength   := int64(60)
	featureSize      := int64(5)
	lastSequenceData := data[len(data) - int(sequenceLength):]

	// DEBUG
	fmt.Println("[ -- fmt.Println(lastSequenceData) -- ]")
	fmt.Println(lastSequenceData)
	fmt.Println()

	fmt.Println("[ -- fmt.Println(len(lastSequenceData)) -- ]")
	fmt.Println(len(lastSequenceData))
	fmt.Println()


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetInputData := time.Now()

	inputData  := make([]float32, sequenceLength * featureSize)
	for idxRow, row := range lastSequenceData {
		for idxFeature := 1; idxFeature <= int(featureSize); idxFeature++ {
			valueFeature, ok := row[idxFeature].(float32)
			if !ok {log.Fatalf(
				"Expected float32 at row %d col %d, got %T", 
				idxRow, idxFeature, row[idxFeature],
			)}

			inputData[idxRow*int(featureSize) + (idxFeature-1)] = valueFeature
		}
	}

	elapsedTimeGetInputData := time.Since(startTimeGetInputData)
	fmt.Printf(
		"[ Time ] Get Input Data : %.9f s\n\n", 
		elapsedTimeGetInputData.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(inputData) -- ]")
	fmt.Println(inputData)
	fmt.Println()

	fmt.Println("[ -- fmt.Println(len(inputData)) -- ]")
	fmt.Println(len(inputData))
	fmt.Println()


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetInputTensor := time.Now()

	inputShape := onnxruntime.NewShape(1, sequenceLength, featureSize)
	inputTensor, err := onnxruntime.NewTensor(inputShape, inputData)
	if err != nil { log.Fatalf("Error creating input tensor: %v", err) }

	elapsedTimeGetInputTensor := time.Since(startTimeGetInputTensor)
	fmt.Printf(
		"[ Time ] Get Input Tensor : %.9f s\n\n", 
		elapsedTimeGetInputTensor.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(inputTensor) -- ]")
	fmt.Println(inputTensor)
	fmt.Println()


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetOutputTensor := time.Now()

	outputShape := onnxruntime.NewShape(1, 1)
	outputTensor, err := onnxruntime.NewEmptyTensor[float32](outputShape)
	if err != nil { log.Fatalf("Error creating output tensor: %v", err) }

	elapsedTimeGetOutputTensor := time.Since(startTimeGetOutputTensor)
	fmt.Printf(
		"[ Time ] Get Output Tensor : %.9f s\n\n", 
		elapsedTimeGetOutputTensor.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(outputTensor) -- ]")
	fmt.Println(outputTensor) 
	fmt.Println()


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetSession := time.Now()
	helpers.MemoryUsage("Before Session")

	session, err := onnxruntime.NewAdvancedSession(
		fmt.Sprintf("./models/%s.onnx", issuer),
		[]string{"input"},   []string{"output"},
		[]onnxruntime.ArbitraryTensor{inputTensor},
		[]onnxruntime.ArbitraryTensor{outputTensor}, nil,
	)
	if err != nil { log.Fatalf("Error initializing ONNX session: %v", err) }

	_ = session
	helpers.MemoryUsage("After Session")

	elapsedTimeGetSession := time.Since(startTimeGetSession)
	fmt.Printf(
		"[ Time ] Get Session : %.9f s\n\n", 
		elapsedTimeGetSession.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(session) -- ]")
	fmt.Println(session) 
	fmt.Println()


	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------
	startTimeGetPredictedData := time.Now()

	var predicted []StockPrice
	lastDate, _ := time.Parse("2006-01-02", actuals[len(actuals)-1].Date)
	for i := 0; i < days; i++ {
		if err := session.Run(); err != nil {
			log.Fatalf("Error running model: %v", err)
		}

		predictedClose := outputTensor.GetData()[0]
		denormPrice := Denormalization(predictedClose, scalers.MinValue["Close"], scalers.MaxValue["Close"])

		lastDate = lastDate.AddDate(0, 0, 1)
		predicted = append(predicted, StockPrice{
			Date:  lastDate.Format("2006-01-02"),
			Price: float32(denormPrice),
		})

		copy(inputData, inputData[int(featureSize):])
		inputData[len(inputData)-1] = predictedClose
	}

	elapsedTimeGetPredictedData := time.Since(startTimeGetPredictedData)
	fmt.Printf(
		"[ Time ] Get Predicted Data : %.9f s\n\n", 
		elapsedTimeGetPredictedData.Seconds(),
	)
	// -----------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------

	// DEBUG
	fmt.Println("[ -- fmt.Println(predicted) -- ]")
	fmt.Println(predicted) 
	fmt.Println()

	fmt.Println("[ -- fmt.Println(len(predicted)) -- ]")
	fmt.Println(len(predicted)) 
	fmt.Println()

	return actuals, predicted, nil
}
