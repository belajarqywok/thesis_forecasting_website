package middlewares

import "github.com/gofiber/fiber/v2"

func NoCacheMiddleware(c *fiber.Ctx) error {
	c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
  c.Set("Pragma", "no-cache")
  c.Set("Expires", "0")

  return c.Next()	
}
