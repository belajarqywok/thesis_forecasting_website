package main

import (
  "fmt"
  "log"

  "os"
	"os/signal"
	"syscall"

  "github.com/joho/godotenv"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"

  helpers      "thesis_forecasting_website/helpers"
  handlers     "thesis_forecasting_website/handlers"
  inferences   "thesis_forecasting_website/inferences"
  middlewares  "thesis_forecasting_website/middlewares"
)

func main() {
  _ = godotenv.Load()

  frontend_engine := html.New("./views", ".html")
  frontend_engine.AddFunc("toJSON", helpers.ToJSON)

  forecasting_service := fiber.New(fiber.Config{
    Views: frontend_engine,
  })

  forecasting_service.Use(middlewares.LoggingMiddleware)
	forecasting_service.Use(helmet.New())
  forecasting_service.Use(middlewares.NoCacheMiddleware)

  forecasting_service.Static("/public", "./public_dist")

  forecasting_service.Get("/", handlers.IssuerHandler)
  forecasting_service.Get("/infographic", handlers.InfographicHandler)
  forecasting_service.Post("/prediction", inferences.StockPredictionHandler)

  host := os.Getenv("FORECASTING_SERVICE_HOST")
	port := os.Getenv("FORECASTING_SERVICE_PORT")

	go func() {
		if err := forecasting_service.Listen(fmt.Sprintf("%s:%s", host, port))
    err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down gracefully...")
	if err := forecasting_service.Shutdown()
  err != nil {
		log.Fatalf("error shutting down: %v", err)
	}
}
