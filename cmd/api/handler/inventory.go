package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/middleware"
	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"

	"github.com/labstack/echo/v4"
)

func MakeInventoryHandlers(e *echo.Group, service inventory.UseCase) {
	e.POST("/inventories", createInventory(service), middleware.AdminRequired)
	e.GET("/inventories", readInventory(service))
	e.GET("/inventories/:id", getInventory(service))
	e.PUT("/inventories/:id", updateInventory(service), middleware.AdminRequired)
	e.DELETE("/inventories/:id", deleteInventory(service), middleware.AdminRequired)
}

func createInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input presenter.Inventory
		if err := c.Bind(&input); err != nil {
			return echo.ErrBadRequest
		}

		pUIID, err := entity.StringToID(input.ProductID)
		if err != nil {
			return echo.ErrBadRequest
		}

		umUIID, err := entity.StringToID(input.UnitOfMeasureID)
		if err != nil {
			return echo.ErrBadRequest
		}

		uUIID, err := entity.StringToID(input.UserID)
		if err != nil {
			return echo.ErrBadRequest
		}

		id, err := service.CreateInventory(uUIID, pUIID, input.Quantity, umUIID)
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

func getInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		data, err := service.GetInventoryByID(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, &presenter.Inventory{
			ID:              data.ID.String(),
			UserID:          data.UserID.String(),
			ProductID:       data.ProductID.String(),
			Quantity:        data.Quantity,
			UnitOfMeasureID: data.UnitOfMeasureID.String(),
		})
	}
}

func readInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.ListInventory()

		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		var out []*presenter.Inventory
		for _, d := range data {
			out = append(out, &presenter.Inventory{
				ID:              d.ID.String(),
				UserID:          d.UserID.String(),
				ProductID:       d.ProductID.String(),
				Quantity:        d.Quantity,
				UnitOfMeasureID: d.UnitOfMeasureID.String(),
			})
		}

		return c.JSON(http.StatusOK, out)
	}
}

func updateInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		var input presenter.Inventory
		if err := c.Bind(&input); err != nil {
			return echo.ErrInternalServerError
		}

		pUIID, err := entity.StringToID(input.ProductID)
		if err != nil {
			return echo.ErrBadRequest
		}

		umUIID, err := entity.StringToID(input.UnitOfMeasureID)
		if err != nil {
			return echo.ErrBadRequest
		}

		uUIID, err := entity.StringToID(input.UserID)
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.UpdateInventory(&entity.Inventory{
			ID:              idUUID,
			UserID:          uUIID,
			ProductID:       pUIID,
			Quantity:        input.Quantity,
			UnitOfMeasureID: umUIID,
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

func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		idUUID, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		err = service.DeleteInventory(idUUID)
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
