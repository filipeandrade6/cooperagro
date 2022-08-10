package echo

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"

	"github.com/labstack/echo/v4"
)

func MakeBaseProductHandlers(e *echo.Echo, service baseproduct.UseCase) {
	e.POST("/baseproducts", createBaseProduct(service))
	e.GET("/baseproducts", readBaseProduct(service))
	e.GET("/baseproducts/:id", getBaseProduct(service))
	e.PUT("/baseproducts/:id", updateBaseProduct(service))
	e.DELETE("/baseproducts/:id", deleteBaseProduct(service))
}

func createBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoCreateBaseProduct

		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "could not get values from the request"},
			)
		}

		id, err := service.CreateBaseProduct(input.Name)
		if errors.Is(entity.ErrEntityAlreadyExists, err) {
			return c.JSON(
				http.StatusConflict,
				presenter.Response{Status: "base product already exists"},
			)
		}
		if errors.Is(entity.ErrInvalidEntity, err) {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "invalid parameters"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: err.Error()},
			)
		}

		return c.JSON(
			http.StatusCreated,
			presenter.Response{
				ID:     id.String(),
				Status: "base product created",
			},
		)
	}
}

func getBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "invalid id"},
			)
		}

		data, err := service.GetBaseProductByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return c.JSON(
				http.StatusNotFound,
				presenter.Response{Status: "base product not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: err.Error()}, // TODO - não expor o erro ao usuŕio?
			)
		}

		return c.JSON(http.StatusOK, &presenter.BaseProduct{
			ID:   data.ID,
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
			return c.JSON(
				http.StatusNotFound,
				presenter.Response{Status: "base products not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: err.Error()},
			)
		}

		var out []*presenter.BaseProduct
		for _, d := range data {
			out = append(out, &presenter.BaseProduct{
				ID:   d.ID,
				Name: d.Name,
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if id == "" {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "invalid id"},
			)
		}

		var input presenter.EchoUpdateBaseProduct
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: "could not get values from the request"},
			)
		}

		if err := service.UpdateBaseProduct(&entity.BaseProduct{
			ID:   idUUID,
			Name: input.Name,
		}); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: err.Error()},
			)
		}

		return c.JSON(http.StatusOK, presenter.Response{Status: "base product udpated"})
	}
}

func deleteBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if id == "" {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				presenter.Response{Status: "invalid id"},
			)
		}

		if err := service.DeleteBaseProduct(idUUID); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				presenter.Response{Status: err.Error()},
			)
		}

		return c.JSON(http.StatusOK, presenter.Response{Status: "base product deleted"})
	}
}
