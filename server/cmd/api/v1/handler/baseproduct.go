package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/v1/middleware"
	"github.com/filipeandrade6/cooperagro/cmd/api/v1/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"

	"github.com/labstack/echo/v4"
)

func MakeBaseProductHandlers(e *echo.Group, service baseproduct.UseCase) {
	e.POST("/baseproducts", createBaseProduct(service), middleware.AdminRequired)
	e.GET("/baseproducts", readBaseProduct(service))
	e.GET("/baseproducts/:id", getBaseProduct(service))
	e.PUT("/baseproducts/:id", updateBaseProduct(service), middleware.AdminRequired)
	e.DELETE("/baseproducts/:id", deleteBaseProduct(service), middleware.AdminRequired)
}

func createBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.BaseProduct
		if err := c.Bind(&input); err != nil {
			return echo.ErrBadRequest
		}

		id, err := service.CreateBaseProduct(input.Name)
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

func getBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		data, err := service.GetBaseProductByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, &presenter.BaseProduct{
			ID:   data.ID.String(),
			Name: data.Name,
		})
	}
}

func readBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*entity.BaseProduct
		var err error

		name := c.QueryParam("name")
		if name != "" {
			data, err = service.SearchBaseProduct(name)
		} else {
			data, err = service.ListBaseProduct()
		}

		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		var out []*presenter.BaseProduct
		for _, d := range data {
			out = append(out, &presenter.BaseProduct{
				ID:   d.ID.String(),
				Name: d.Name,
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		var input presenter.BaseProduct
		if err := c.Bind(&input); err != nil {
			return echo.ErrInternalServerError
		}

		err = service.UpdateBaseProduct(&entity.BaseProduct{
			ID:   idUUID,
			Name: input.Name,
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

func deleteBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.DeleteBaseProduct(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
