package echo

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
// 	"github.com/labstack/echo/v4"
// )

// func MakeInventoryHandlers(e *echo.Echo, service inventory.UseCase) {
// 	e.POST("/inventory/", createInventory(service))
// 	e.GET("/inventory/", readInventory(service))
// 	e.PUT("/inventory/", updateInventory(service))
// 	e.DELETE("/inventory/", deleteInventory(service))
// }

// func createInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoCreateInventory

// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				presenter.Response{Status: "could not get values from the request"},
// 			)
// 		}

// 		id, err := service.CreateInventory(input.Name)
// 		if errors.Is(entity.ErrEntityAlreadyExists, err) {
// 			return c.JSON(
// 				http.StatusConflict,
// 				presenter.Response{Status: "base product already exists"},
// 			)
// 		}
// 		if errors.Is(entity.ErrInvalidEntity, err) {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				presenter.Response{Status: "invalid parameters"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: err.Error()},
// 			)
// 		}

// 		return c.JSON(
// 			http.StatusCreated,
// 			presenter.Response{
// 				ID:     id.String(),
// 				Status: "base product created",
// 			},
// 		)
// 	}
// }

// func readInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoInventory
// 		var datas []*entity.Inventory
// 		var err error

// 		if err = c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				presenter.Response{Status: "could not get values from the request"},
// 			)
// 		}

// 		switch {
// 		case input.ID != "":
// 			idUUID, err := entity.StringToID(input.ID)
// 			if err != nil {
// 				return c.JSON(
// 					http.StatusBadRequest,
// 					presenter.Response{Status: "invalid id"},
// 				)
// 			}

// 			data, err := service.GetInventoryByID(idUUID)
// 			if errors.Is(err, entity.ErrNotFound) {
// 				return c.JSON(
// 					http.StatusNotFound,
// 					presenter.Response{Status: "base product not found"},
// 				)
// 			}
// 			if err != nil {
// 				return c.JSON(
// 					http.StatusInternalServerError,
// 					presenter.Response{Status: err.Error()}, // TODO - não expor o erro ao usuŕio?
// 				)
// 			}

// 			return c.JSON(http.StatusOK, &presenter.Inventory{
// 				ID:   data.ID,
// 				Name: data.Name,
// 			})

// 		case input.Name == "":
// 			datas, err = service.ListInventory()

// 		default:
// 			datas, err = service.SearchInventory(input.Name)

// 		}

// 		if errors.Is(err, entity.ErrNotFound) {
// 			return c.JSON(
// 				http.StatusNotFound,
// 				presenter.Response{Status: "base product not found"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: err.Error()},
// 			)
// 		}

// 		var toJ []*presenter.Inventory
// 		for _, d := range datas {
// 			toJ = append(toJ, &presenter.Inventory{
// 				ID:   d.ID,
// 				Name: d.Name,
// 			})
// 		}

// 		return c.JSON(http.StatusOK, toJ)
// 	}
// }

// func updateInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoUpdateInventory
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: "could not get values from the request"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(input.ID)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				presenter.Response{Status: "invalid id"},
// 			)
// 		}

// 		if err := service.UpdateInventory(&entity.Inventory{
// 			ID:   idUUID,
// 			Name: input.Name,
// 		}); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, presenter.Response{Status: "base product udpated"})
// 	}
// }

// func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoDeleteInventory
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: "could not get values from the request"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(input.ID)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				presenter.Response{Status: "invalid id"},
// 			)
// 		}

// 		if err := service.DeleteInventory(idUUID); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				presenter.Response{Status: err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, presenter.Response{Status: "base product deleted"})
// 	}
// }
