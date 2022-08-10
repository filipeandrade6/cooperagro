package echo

import (
	"errors"
	"net/http"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/filipeandrade6/cooperagro/infra/auth"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func MakeAuthHandlers(e *echo.Group, service user.UseCase) {
	e.POST("/login", login(service))
}

func login(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// TODO remover isso aqui
		id, err := entity.StringToID(username)
		if err != nil {
			return echo.ErrUnauthorized // TODO tipo de erro errado
		}

		user, err := service.GetUserByID(id)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}

		if err = user.ValidatePassword(password); err != nil {
			return echo.ErrUnauthorized
		}

		// Set custom claims
		claims := &auth.JWTCustomClaims{
			ID:    username,
			Roles: user.Roles,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}
