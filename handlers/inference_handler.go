package handlers

import (
	"github.com/gofiber/fiber/v2"
	inferences  "thesis_forecasting_website/inferences"
)

type StockRequest struct {
	Issuer string `json:"issuer"`
	Days   int    `json:"days"`
}

type StockResponse struct {
	Actuals    []inferences.StockPrice `json:"actuals"`
	Prediction []inferences.StockPrice `json:"predictions"`
}

func InferenceHandler(context *fiber.Ctx) error {
	request := new(StockRequest)
	if err  := context.BodyParser(request); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Days <= 0 { request.Days = 7 }
	if request.Days > 60 { request.Days = 60 }

	// actuals, predicted, err := inferences.StockPredictionDebug(
	// 	request.Issuer, request.Days,
	// )

	actuals, predicted, err := inferences.StockPrediction(
		request.Issuer, request.Days,
	)

	if err != nil {
    return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Internal server error",
    })
	}

	return context.Status(fiber.StatusOK).JSON(StockResponse{
		Actuals:    actuals,
		Prediction: predicted,
	})
}