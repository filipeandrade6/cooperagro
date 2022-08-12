package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/v1/middleware"
	"github.com/filipeandrade6/cooperagro/cmd/api/v1/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/labstack/echo/v4"
)

func MakeUserHandlers(e *echo.Group, service user.UseCase) {
	e.POST("/users", createUser(service), middleware.AdminRequired)
	e.GET("/users", readUser(service), middleware.AdminRequired) // TODO alterar isso aqui (buyer podem ver producers?)
	e.GET("/users/:id", getUser(service), middleware.AdminRequired)
	e.PUT("/users/:id", updateUser(service), middleware.AdminRequired)
	e.DELETE("/users/:id", deleteUser(service), middleware.AdminRequired)
}

func createUser(service user.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.User
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

		return c.JSON(http.StatusOK, &presenter.User{
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

		var out []*presenter.User
		for _, d := range data {
			out = append(out, &presenter.User{
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

		var input presenter.User
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
		case errors.Is(entity.ErrEntityAlreadyExists, err):
			return c.NoContent(http.StatusConflict)

		case errors.Is(entity.ErrInvalidEntity, err):
			return echo.ErrBadRequest

		case errors.Is(entity.ErrNotFound, err):
			return echo.ErrNotFound

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
