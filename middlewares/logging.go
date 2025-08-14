package middlewares

import (
	"log"
	"time"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start_time  := time.Now()
	err         := c.Next()
	status_code := c.Response().StatusCode()
	latency     := time.Since(start_time)

	log.Printf("[%s] %s %s %s | Status: %d %s | IP: %s | UA: %s | Latency: %v",
		time.Now().Format("2006-01-02 15:04:05"),
		c.Method(), c.Path(), string(c.Request().URI().QueryString()),
		status_code, http.StatusText(status_code),
		c.IP(), c.Get("User-Agent"), latency,
	)

	return err
}
