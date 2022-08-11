package middleware

import (
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ClaimsContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		c.Set("claims", user.Claims.(*auth.Claims))

		return next(c)
	}
}

func AdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("claims").(*auth.Claims)
		if !claims.Authorized(entity.RoleAdmin) {
			return echo.ErrForbidden
		}
		return next(c)
	}
}

func ProducerRequired(next echo.HandlerFunc) echo.HandlerFunc { // TODO deletar se não tiver uso
	return func(c echo.Context) error {
		claims := c.Get("claims").(*auth.Claims)
		if !claims.Authorized(entity.RoleProducer) {
			return echo.ErrForbidden
		}
		return next(c)
	}
}

func BuyerRequired(next echo.HandlerFunc) echo.HandlerFunc { // TODO deletar se não tiver uso
	return func(c echo.Context) error {
		claims := c.Get("claims").(*auth.Claims)
		if !claims.Authorized(entity.RoleBuyer) {
			return echo.ErrForbidden
		}
		return next(c)
	}
}
