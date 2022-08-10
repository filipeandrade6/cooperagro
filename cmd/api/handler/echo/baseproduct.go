package echo

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"

	"github.com/labstack/echo/v4"
)

func MakeBaseProductHandlers(e *echo.Group, service baseproduct.UseCase) {
	e.POST("/baseproducts", createBaseProduct(service))
	e.GET("/baseproducts", readBaseProduct(service))
	e.GET("/baseproducts/:id", getBaseProduct(service))
	e.PUT("/baseproducts/:id", updateBaseProduct(service))
	e.DELETE("/baseproducts/:id", deleteBaseProduct(service))
}

func createBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.BaseProduct

		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "could not get values from the request"},
			)
		}

		id, err := service.CreateBaseProduct(input.Name)
		if errors.Is(entity.ErrEntityAlreadyExists, err) {
			return c.JSON(
				http.StatusConflict,
				echo.Map{"status": "base product already exists"},
			)
		}
		if errors.Is(entity.ErrInvalidEntity, err) {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "invalid parameters"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		return c.JSON(
			http.StatusCreated,
			echo.Map{"id": id.String()},
		)
	}
}

func getBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "invalid id"},
			)
		}

		data, err := service.GetBaseProductByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return c.JSON(
				http.StatusNotFound,
				echo.Map{"status": "base product not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()}, // TODO - não expor o erro ao usuŕio?
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
				echo.Map{"status": "base products not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
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
				echo.Map{"status": "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "invalid id"},
			)
		}

		var input presenter.BaseProduct
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": "could not get values from the request"},
			)
		}

		if err := service.UpdateBaseProduct(&entity.BaseProduct{
			ID:   idUUID,
			Name: input.Name,
		}); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		return c.JSON(http.StatusOK, echo.Map{"status": "base product udpated"})
	}
}

func deleteBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if id == "" {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "empty id"},
			)
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "invalid id"},
			)
		}

		if err := service.DeleteBaseProduct(idUUID); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		return c.JSON(http.StatusOK, echo.Map{"status": "base product deleted"})
	}
}
