package echo

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"

	"github.com/labstack/echo/v4"
)

// TODO juntar em um só get e list

func MakeBaseProductHandlers(e *echo.Echo, service baseproduct.UseCase) {
	// e.GET("/baseproduct/:id", getBaseProductByID(service))
	e.GET("/baseproduct/", getAndlistBaseProduct(service))
	e.POST("/baseproduct/", createBaseProduct(service))
	e.PUT("/baseproduct/", updateBaseProduct(service))
	e.DELETE("/baseproduct/", deleteBaseProduct(service))
}

// func getBaseProductByID(service baseproduct.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, err := entity.StringToID(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, "invalid id")
// 		}

// 		data, err := service.GetBaseProductByID(id)
// 		if errors.Is(err, entity.ErrNotFound) {
// 			return c.JSON(http.StatusNotFound, "not found")
// 		}
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error()) // TODO - não expor o erro ao usuŕio?
// 		}

// 		return c.JSON(http.StatusOK, &presenter.BaseProduct{
// 			ID:   data.ID,
// 			Name: data.Name,
// 		})
// 	}
// }

func getAndlistBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoBaseProduct
		var datas []*entity.BaseProduct
		var err error

		if err = c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request")
		}

		switch {
		case input.ID != "":
			idUUID, err := entity.StringToID(input.ID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, "invalid id")
			}
			data, err := service.GetBaseProductByID(idUUID)
			if errors.Is(err, entity.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error()) // TODO - não expor o erro ao usuŕio?
			}

			return c.JSON(http.StatusOK, &presenter.BaseProduct{
				ID:   data.ID,
				Name: data.Name,
			})

		case input.Name == "":
			datas, err = service.ListBaseProduct()
		default:
			datas, err = service.SearchBaseProduct(input.Name)
		}

		if errors.Is(err, entity.ErrNotFound) {
			return c.JSON(http.StatusNotFound, "not found")
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var toJ []*presenter.BaseProduct
		for _, d := range datas {
			toJ = append(toJ, &presenter.BaseProduct{
				ID:   d.ID,
				Name: d.Name,
			})
		}

		return c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoCreateBaseProduct
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				"could not get values from the request",
			)
		}

		id, err := service.CreateBaseProduct(input.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, id)
		// Se der erro de marshalling no JSON?
	}
}

func updateBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoUpdateBaseProduct
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				"could not get values from the request",
			)
		}

		if input.ID == "" {
			return c.JSON(http.StatusBadRequest, "empty id")
		}

		idUUID, err := entity.StringToID(input.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}

		if err := service.UpdateBaseProduct(&entity.BaseProduct{
			ID:   idUUID,
			Name: input.Name,
		}); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "base product udpated")
	}
}

func deleteBaseProduct(service baseproduct.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoDeleteBaseProduct
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				"could not get values from the request",
			)
		}

		if input.ID == "" {
			return c.JSON(http.StatusBadRequest, "empty id")
		}

		idUUID, err := entity.StringToID(input.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}

		if err := service.DeleteBaseProduct(idUUID); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "base product deleted")
	}
}
