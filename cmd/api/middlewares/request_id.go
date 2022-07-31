package middlewares

import (
	"github.com/filipeandrade6/cooperagro-go/domain"
	"github.com/gofiber/fiber/v2"
)

// HandleRequestID will try to read a `request-id`
// key from the request headers and if it is not available
// generate a random one and put it in the request context.
func HandleRequestID() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("request-id")
		if requestID == "" {
			requestID = domain.GenerateRequestID()
		}

		c.Locals(domain.RequestIDKey, requestID)
		return c.Next()
	}
}
