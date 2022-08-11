package echo

import (
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func NeedUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.Claims)

		c.Set("userID", claims.UserID)

		return next(c)
	}
}
