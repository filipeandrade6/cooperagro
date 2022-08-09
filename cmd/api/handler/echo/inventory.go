package echo

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
// 	"github.com/gin-gonic/gin"
// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/v4"
// )

// func MakeInventoryHandlers(e *echo.Echo, service inventory.UseCase) {
// 	e.GET("/inventory/:id", getInventoryByID(service))
// 	e.GET("/inventory", listInventory(service))
// 	e.POST("/inventory", createInventory(service))
// 	e.PUT("/inventory/:id", updateInventory(service))
// 	e.DELETE("/inventory/:id", deleteInventory(service))
// }

// func getInventoryByID(service inventory.UseCase) echo.HandlerFunc {
// 	return func(e *echo.Context) error {
// 		errorMessage := "error reading inventory"

// 		id, err := entity.StringToID(e.Param("id"))
// 		if err != nil {
// 			e.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
// 			return
// 		}

// 		data, err := service.GetInventoryByID(id)

// 		if err != nil && !errors.Is(err, entity.ErrNotFound) {
// 			e.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
// 			return
// 		}

// 		if data == nil {
// 			e.JSON(http.StatusNotFound, gin.H{"status": "not found"})
// 			return
// 		}

// 		e.JSON(http.StatusOK, &presenter.Inventory{
// 			ID:              data.ID,
// 			UserID:          data.UserID,
// 			ProductID:       data.ProductID,
// 			Quantity:        data.Quantity,
// 			UnitOfMeasureID: data.UnitOfMeasureID,
// 		})

// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func listInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(e *echo.Context) {
// 		errorMessage := "error reading inventory"

// 		data, err := service.ListInventory()

// 		if err != nil && !errors.Is(err, entity.ErrNotFound) {
// 			e.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
// 			return
// 		}

// 		if data == nil {
// 			e.JSON(http.StatusNotFound, gin.H{"status": "not found"})
// 			return
// 		}

// 		var toJ []*presenter.Inventory
// 		for _, d := range data {
// 			toJ = append(toJ, &presenter.Inventory{
// 				ID:              d.ID,
// 				UserID:          d.UserID,
// 				ProductID:       d.ProductID,
// 				Quantity:        d.Quantity,
// 				UnitOfMeasureID: d.UnitOfMeasureID,
// 			})
// 		}
// 		e.JSON(http.StatusOK, toJ)

// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func createInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(e *echo.Context) {
// 		var input presenter.CreateInventory
// 		if err := e.ShouldBindJSON(&input); err != nil {
// 			e.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
// 			return
// 		}

// 		id, err := service.CreateInventory(
// 			input.UserID,
// 			input.ProductID,
// 			input.Quantity,
// 			input.UnitOfMeasureID,
// 		)
// 		if err != nil {
// 			e.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		e.JSON(http.StatusCreated, gin.H{"id": id})
// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func updateInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(e *echo.Context) {
// 		id := e.Param("id")

// 		if id == "" {
// 			e.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
// 			return
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			e.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 			return
// 		}

// 		var input presenter.UpdateInventory
// 		if err := e.ShouldBindJSON(&input); err != nil {
// 			e.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := service.UpdateInventory(&entity.Inventory{
// 			ID:              idUUID,
// 			UserID:          input.UserID,
// 			ProductID:       input.ProductID,
// 			Quantity:        input.Quantity,
// 			UnitOfMeasureID: input.UnitOfMeasureID,
// 		}); err != nil {
// 			e.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		e.JSON(http.StatusOK, gin.H{"status": "inventory udpated"})
// 	}
// }

// func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(e *echo.Context) {
// 		id := e.Param("id")
// 		if id == "" {
// 			e.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
// 			return
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			e.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 			return
// 		}

// 		if err := service.DeleteInventory(idUUID); err != nil {
// 			e.JSON(http.StatusInternalServerError, gin.H{"error": "deleting inventory"})
// 			return
// 		}

// 		e.JSON(http.StatusOK, gin.H{"status": "inventory deleted"})
// 	}
// }
