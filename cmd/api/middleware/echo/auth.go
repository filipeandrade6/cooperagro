package echo

import (
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roles := c.Get("roles").([]string)

		for _, want := range entity.Roles {
			for _, has := range roles {
				if has == want {
					return next(c)
				}
			}
		}

		return echo.ErrForbidden
	}
}

func Logado(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.Claims)

		c.Set("userID", claims.UserID)
		c.Set("roles", claims.Roles)

		return next(c)
	}
}
