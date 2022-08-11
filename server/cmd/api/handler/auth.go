package handler

import (
	"errors"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/filipeandrade6/cooperagro/infra/auth"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func MakeAuthHandlers(e *echo.Echo, service user.UseCase) {
	e.GET("/login/", loginPage)
	e.POST("/login", login(service))
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

// TODO remover abaixo

func loginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
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
		claims := &auth.Claims{
			UserID: username,
			Roles:  user.Roles,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
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
