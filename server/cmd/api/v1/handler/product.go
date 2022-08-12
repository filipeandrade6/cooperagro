package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/v1/middleware"
	"github.com/filipeandrade6/cooperagro/cmd/api/v1/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
	"github.com/labstack/echo/v4"
)

func MakeProductHandlers(e *echo.Group, service product.UseCase) {
	e.POST("/products", createProduct(service), middleware.AdminRequired)
	e.GET("/products", readProduct(service))
	e.GET("/products/:id", getProduct(service))
	e.PUT("/products/:id", updateProduct(service), middleware.AdminRequired)
	e.DELETE("/products/:id", deleteProduct(service), middleware.AdminRequired)
}

func createProduct(service product.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.Product
		if err := c.Bind(&input); err != nil {
			return echo.ErrBadRequest
		}

		bpUIID, err := entity.StringToID(input.BaseProductID)
		if err != nil {
			return echo.ErrBadRequest
		}

		id, err := service.CreateProduct(input.Name, bpUIID)
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

		return c.JSON(
			http.StatusCreated,
			echo.Map{"id": id.String()},
		)
	}
}

func getProduct(service product.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		data, err := service.GetProductByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, &presenter.Product{
			ID:            data.ID.String(),
			Name:          data.Name,
			BaseProductID: data.BaseProductID.String(),
		})
	}
}

func readProduct(service product.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*entity.Product
		var err error

		name := c.QueryParam("name")
		if name != "" {
			data, err = service.SearchProduct(name)
		} else {
			data, err = service.ListProduct()
		}

		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		var out []*presenter.Product
		for _, d := range data {
			out = append(out, &presenter.Product{
				ID:            d.ID.String(),
				Name:          d.Name,
				BaseProductID: d.BaseProductID.String(),
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateProduct(service product.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		var input presenter.Product
		if err := c.Bind(&input); err != nil {
			return echo.ErrInternalServerError
		}

		bpUIID, err := entity.StringToID(input.BaseProductID)
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.UpdateProduct(&entity.Product{
			ID:            idUUID,
			Name:          input.Name,
			BaseProductID: bpUIID,
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

func deleteProduct(service product.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.DeleteProduct(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
