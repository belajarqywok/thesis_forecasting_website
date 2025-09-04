package inferences

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	loaders     "thesis_forecasting_website/loaders"
	onnxruntime "github.com/belajarqywok/onnxruntime_go"
	"sync"
)

var once sync.Once
var isInit bool

func denormalization(data, minValue, maxValue float32) float32 {
	return (data * (maxValue - minValue)) + minValue
}

type StockRequest struct {
	Issuer string `json:"issuer"`
	Days   int    `json:"days"`
}

type StockPrice struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

type StockResponse struct {
	Actuals    []StockPrice `json:"actuals"`
	Prediction []StockPrice `json:"prediction"`
}

func StockPredictionHandler(c *fiber.Ctx) error {
	req := new(StockRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Days <= 0 { req.Days = 7 }

	once.Do(func() {
		onnxruntime.SetSharedLibraryPath("./onnxruntime-linux-x64-1.21.0/lib/libonnxruntime.so")
		if err := onnxruntime.InitializeEnvironment(); err != nil {
			log.Fatal("Error initializing ONNX runtime: ", err)
		}

		isInit = true
	})

	if !isInit { log.Fatal("ONNX runtime not initialized") }

	// load dataset
	dataset_csv_path := fmt.Sprintf("./indonesia_stocks/modeling_datas/%s.csv", req.Issuer)
	data, err := loaders.DatasetLoader(dataset_csv_path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error loading CSV",
		})
	}

	// load scaler
	minmax_json_path := fmt.Sprintf("./indonesia_stocks/min_max/%s.json", req.Issuer)
	scalers, err := loaders.ScalersLoader(minmax_json_path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error loading scalers",
		})
	}

	n := 30
	if len(data) < n { n = len(data) }
	lastDataActual := data[len(data) - n:]

	var actuals []StockPrice
	for _, row := range lastDataActual {
		date := row[0].(string)
		closeVal := row[1].(float32)

		closePrice := denormalization(
			closeVal,
			scalers.MinValue["Close"],
			scalers.MaxValue["Close"],
		)

		actuals = append(actuals, StockPrice{
			Date:  date,
			Price: float64(closePrice),
		})
	}


	// Prepare input for model
	sequenceLength := int64(60)
	featureSize    := int64(5)
	lastData       := data[len(data) - int(sequenceLength):]

	// inputData  := make([]float32, sequenceLength * featureSize)
	// for i, row := range lastData {
	// 	for j := 1; j <= int(featureSize); j++ {
	// 		val, ok := row[j].(float32)
	// 		if !ok {
	// 			log.Fatalf("Expected float32 at row %d col %d, got %T", i, j, row[j])
	// 		}

	// 		inputData[i*int(featureSize) + (j-1)] = val
	// 	}
	// }

	inputData := make([]float32, sequenceLength*featureSize)
	for i, row := range lastData {
		copy(inputData[i*int(featureSize) : (i+1)*int(featureSize)], row[1:])
	}

	inputShape := onnxruntime.NewShape(1, sequenceLength, featureSize)
	inputTensor, err := onnxruntime.NewTensor(inputShape, inputData)
	if err != nil { log.Fatalf("Error creating input tensor: %v", err) }

	outputShape := onnxruntime.NewShape(1, 1)
	outputTensor, err := onnxruntime.NewEmptyTensor[float32](outputShape)
	if err != nil { log.Fatalf("Error creating output tensor: %v", err) }

	model_onnx_path := fmt.Sprintf("./models/%s.onnx", req.Issuer)
	session, err := onnxruntime.NewAdvancedSession(
		model_onnx_path,
		[]string{"input"}, []string{"output"},
		[]onnxruntime.ArbitraryTensor{inputTensor},
		[]onnxruntime.ArbitraryTensor{outputTensor}, nil,
	)

	if err != nil { log.Fatalf("Error initializing ONNX session: %v", err) }

	// generate predictions
	var predicted []StockPrice
	lastDate, _ := time.Parse("2006-01-02", actuals[len(actuals)-1].Date)

	for i := 0; i < req.Days; i++ {
		if err := session.Run(); err != nil {
			log.Fatalf("Error running model: %v", err)
		}

		predictedClose := outputTensor.GetData()[0]
		denormPrice := denormalization(predictedClose, scalers.MinValue["Close"], scalers.MaxValue["Close"])

		lastDate = lastDate.AddDate(0, 0, 1)
		predicted = append(predicted, StockPrice{
			Date:  lastDate.Format("2006-01-02"),
			Price: float64(denormPrice),
		})

		copy(inputData, inputData[int(featureSize):])
		inputData[len(inputData)-1] = predictedClose
	}

	resp := StockResponse{
		Actuals:    actuals,
		Prediction: predicted,
	}

	return c.JSON(resp)
}
