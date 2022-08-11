package echo

import (
	"errors"
	"net/http"

	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware/echo"
	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/labstack/echo/v4"
)

func MakeUserHandlers(e *echo.Group, service user.UseCase) {
	e.POST("/users", createUser(service), mid.AdminRequired)
	e.GET("/users", readUser(service), mid.AdminRequired) // TODO alterar isso aqui (buyer podem ver producers?)
	e.GET("/users/:id", getUser(service), mid.AdminRequired)
	e.PUT("/users/:id", updateUser(service), mid.AdminRequired)
	e.DELETE("/users/:id", deleteUser(service), mid.AdminRequired)
}

func createUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoUser
		if err := c.Bind(&input); err != nil {
			return echo.ErrBadRequest
		}

		id, err := service.CreateUser(
			input.FirstName,
			input.LastName,
			input.Address,
			input.Phone,
			input.Email,
			input.Latitude,
			input.Longitude,
			input.Roles,
			input.Password,
		)
		if errors.Is(entity.ErrEntityAlreadyExists, err) {
			return c.NoContent(http.StatusConflict)
		}
		if errors.Is(entity.ErrInvalidEntity, err) {
			return echo.ErrBadRequest
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(
			http.StatusCreated,
			echo.Map{"id": id.String()},
		)
	}
}

func getUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		data, err := service.GetUserByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, &presenter.EchoUser{
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Address:   data.Address,
			Phone:     data.Phone,
			Email:     data.Email,
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
			Roles:     data.Roles,
		})
	}
}

func readUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*entity.User
		var err error

		firstName := c.QueryParam("first_name")
		if firstName != "" {
			data, err = service.SearchUser(firstName)
		} else {
			data, err = service.ListUser()
		}

		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		var out []*presenter.EchoUser
		for _, d := range data {
			out = append(out, &presenter.EchoUser{
				ID:        d.ID.String(),
				FirstName: d.FirstName,
				LastName:  d.LastName,
				Address:   d.Address,
				Phone:     d.Phone,
				Email:     d.Email,
				Latitude:  d.Latitude,
				Longitude: d.Longitude,
				Roles:     d.Roles,
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		var input presenter.EchoUser
		if err := c.Bind(&input); err != nil {
			return echo.ErrInternalServerError
		}

		err = service.UpdateUser(&entity.User{
			ID:        idUUID,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Address:   input.Address,
			Phone:     input.Phone,
			Email:     input.Email,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Roles:     input.Roles,
			Password:  input.Password,
		})
		switch {
		case errors.Is(entity.ErrInvalidEntity, err):
			return echo.ErrBadRequest

		case errors.Is(entity.ErrNotFound, err):
			return echo.ErrNotFound

		case errors.Is(entity.ErrEntityAlreadyExists, err):
			return c.NoContent(http.StatusConflict)

		case err != nil:
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}

func deleteUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.DeleteUser(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
