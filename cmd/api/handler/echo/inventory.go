package echo

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
	"github.com/labstack/echo/v4"
)

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
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		id, err := service.CreateInventory(input.Name)
// 		if errors.Is(entity.ErrEntityAlreadyExists, err) {
// 			return c.JSON(
// 				http.StatusConflict,
// 				echo.Map{"status": "inventory already exists"},
// 			)
// 		}
// 		if errors.Is(entity.ErrInvalidEntity, err) {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid parameters"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(
// 			http.StatusCreated,
// 			presenter.Response{
// 				ID:     id.String(),
// 				Status: "inventory created",
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
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		switch {
// 		case input.ID != "":
// 			idUUID, err := entity.StringToID(input.ID)
// 			if err != nil {
// 				return c.JSON(
// 					http.StatusBadRequest,
// 					echo.Map{"status": "invalid id"},
// 				)
// 			}

// 			data, err := service.GetInventoryByID(idUUID)
// 			if errors.Is(err, entity.ErrNotFound) {
// 				return c.JSON(
// 					http.StatusNotFound,
// 					echo.Map{"status": "inventory not found"},
// 				)
// 			}
// 			if err != nil {
// 				return c.JSON(
// 					http.StatusInternalServerError,
// 					echo.Map{"status": err.Error()}, // TODO - não expor o erro ao usuŕio?
// 				)
// 			}

// 			return c.JSON(http.StatusOK, &presenter.EchoInventory{
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
// 				echo.Map{"status": "inventory not found"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		var toJ []*presenter.EchoInventory
// 		for _, d := range datas {
// 			toJ = append(toJ, &presenter.EchoInventory{
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
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(input.ID)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid id"},
// 			)
// 		}

// 		if err := service.UpdateInventory(&entity.Inventory{
// 			ID:   idUUID,
// 			Name: input.Name,
// 		}); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"status": "inventory udpated"})
// 	}
// }

// func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoDeleteInventory
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(input.ID)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid id"},
// 			)
// 		}

// 		if err := service.DeleteInventory(idUUID); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"status": "inventory deleted"})
// 	}
// }

func MakeInventoryHandlers(e *echo.Group, service inventory.UseCase) {
	e.POST("/inventories", createInventory(service))
	e.GET("/inventories", readInventory(service))
	e.GET("/inventories/:id", getInventory(service))
	e.PUT("/inventories/:id", updateInventory(service))
	e.DELETE("/inventories/:id", deleteInventory(service))
}

func createInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.EchoInventory

		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{"status": "could not get values from the request"},
			)
		}

		userUIID, err := entity.StringToID(input.UserID)
		productUIID, err := entity.StringToID(input.UserID)
		unitOfMeasureUIID, err := entity.StringToID(input.UnitOfMeasureID)

		id, err := service.CreateInventory(
			userUIID,
			productUIID,
			input.Quantity,
			unitOfMeasureUIID,
		)
		if errors.Is(entity.ErrEntityAlreadyExists, err) {
			return c.JSON(
				http.StatusConflict,
				echo.Map{"status": "inventory already exists"},
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

func getInventory(service inventory.UseCase) echo.HandlerFunc {
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

		data, err := service.GetInventoryByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return c.JSON(
				http.StatusNotFound,
				echo.Map{"status": "inventory not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()}, // TODO - não expor o erro ao usuŕio?
			)
		}

		return c.JSON(http.StatusOK, &presenter.EchoInventory{
			ID:   data.ID,
			Name: data.Name,
		})
	}
}

func readInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*entity.Inventory
		var err error

		name := c.QueryParam("name")
		if name != "" {
			data, err = service.SearchInventory(name)
		} else {
			data, err = service.ListInventory()
		}

		if errors.Is(err, entity.ErrNotFound) {
			return c.JSON(
				http.StatusNotFound,
				echo.Map{"status": "inventorys not found"},
			)
		}
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		var out []*presenter.EchoInventory
		for _, d := range data {
			out = append(out, &presenter.EchoInventory{
				ID:   d.ID,
				Name: d.Name,
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateInventory(service inventory.UseCase) echo.HandlerFunc {
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

		var input presenter.EchoInventory
		if err := c.Bind(&input); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": "could not get values from the request"},
			)
		}

		if err := service.UpdateInventory(&entity.Inventory{
			ID:   idUUID,
			Name: input.Name,
		}); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		return c.JSON(http.StatusOK, echo.Map{"status": "inventory udpated"})
	}
}

func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
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

		if err := service.DeleteInventory(idUUID); err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.Map{"status": err.Error()},
			)
		}

		return c.JSON(http.StatusOK, echo.Map{"status": "inventory deleted"})
	}
}
