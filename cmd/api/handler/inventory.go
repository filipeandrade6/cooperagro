package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entities"
	"github.com/filipeandrade6/cooperagro/domain/usecases/inventory"
	"github.com/gin-gonic/gin"
)

func MakeInventoryHandlers(r *gin.Engine, service inventory.UseCase) {
	r.GET("/inventory/:id", getInventoryByID(service))
	r.GET("/inventory", listInventory(service))
	r.POST("/inventory", createInventory(service))
	r.PUT("/inventory/:id", updateInventory(service))
	r.DELETE("/inventory/:id", deleteInventory(service))
}

func getInventoryByID(service inventory.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading inventory"

		id, err := entities.StringToID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		data, err := service.GetInventoryByID(id)

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		c.JSON(http.StatusOK, &presenter.Inventory{
			ID:            data.ID,
			Customer:      data.Customer,
			Product:       data.Product,
			Quantity:      data.Quantity,
			UnitOfMeasure: data.UnitOfMeasure,
		})

		// Se der erro de marshalling no JSON?
	}
}

func listInventory(service inventory.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading inventory"

		data, err := service.ListInventory()

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		var toJ []*presenter.Inventory
		for _, d := range data {
			toJ = append(toJ, &presenter.Inventory{
				ID:            d.ID,
				Customer:      d.Customer,
				Product:       d.Product,
				Quantity:      d.Quantity,
				UnitOfMeasure: d.UnitOfMeasure,
			})
		}
		c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createInventory(service inventory.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input presenter.CreateInventory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
			return
		}

		id, err := service.CreateInventory(
			input.Customer,
			input.Product,
			input.Quantity,
			input.UnitOfMeasure,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
		// Se der erro de marshalling no JSON?
	}
}

func updateInventory(service inventory.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
			return
		}

		idUUID, err := entities.StringToID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var input presenter.UpdateInventory
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.UpdateInventory(&entities.Inventory{
			ID:            idUUID,
			Customer:      input.Customer,
			Product:       input.Product,
			Quantity:      input.Quantity,
			UnitOfMeasure: input.UnitOfMeasure,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "inventory udpated"})
	}
}

func deleteInventory(service inventory.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
			return
		}

		idUUID, err := entities.StringToID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if err := service.DeleteInventory(idUUID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting inventory"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "inventory deleted"})
	}
}
