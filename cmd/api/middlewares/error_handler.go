package middlewares

import (
	"github.com/filipeandrade6/cooperagro/adapters/log"
	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/gofiber/fiber/v2"
)

// HandleError is responsible for converting domain errors to HTTP errors
// simplifying error handling overall.
func HandleError(logger log.Provider) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		req := c.Request()
		status, body := domain.HandleDomainErrAsHTTP(
			c.Context(),
			logger,
			err,
			string(req.Header.Method()),
			string(req.RequestURI()),
		)
		c.Status(status).Send(body)
		return nil
	}
}
