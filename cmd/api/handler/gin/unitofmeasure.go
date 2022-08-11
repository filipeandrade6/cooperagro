package gin

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
// 	"github.com/gin-gonic/gin"
// )

// func MakeUnitOfMeasureHandlers(r *gin.Engine, service unitofmeasure.UseCase) {
// 	r.GET("/unitofmeasure/:id", getUnitOfMeasureByID(service))
// 	r.GET("/unitofmeasure", listUnitOfMeasure(service))
// 	r.POST("/unitofmeasure", createUnitOfMeasure(service))
// 	r.PUT("/unitofmeasure/:id", updateUnitOfMeasure(service))
// 	r.DELETE("/unitofmeasure/:id", deleteUnitOfMeasure(service))
// }

// func getUnitOfMeasureByID(service unitofmeasure.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		errorMessage := "error reading unit of measure"

// 		id, err := entity.StringToID(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
// 			return
// 		}

// 		data, err := service.GetUnitOfMeasureByID(id)

// 		if err != nil && !errors.Is(err, entity.ErrNotFound) {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
// 			return
// 		}

// 		if data == nil {
// 			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, &presenter.UnitOfMeasure{
// 			ID:   data.ID,
// 			Name: data.Name,
// 		})

// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func listUnitOfMeasure(service unitofmeasure.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		errorMessage := "error reading unit of measure"

// 		data, err := service.ListUnitOfMeasure()

// 		if err != nil && !errors.Is(err, entity.ErrNotFound) {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
// 			return
// 		}

// 		if data == nil {
// 			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
// 			return
// 		}

// 		var toJ []*presenter.UnitOfMeasure
// 		for _, d := range data {
// 			toJ = append(toJ, &presenter.UnitOfMeasure{
// 				ID:   d.ID,
// 				Name: d.Name,
// 			})
// 		}
// 		c.JSON(http.StatusOK, toJ)

// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func createUnitOfMeasure(service unitofmeasure.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var input presenter.CreateUnitOfMeasure
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
// 			return
// 		}

// 		id, err := service.CreateUnitOfMeasure(input.Name)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "creating unit of measure"})
// 			return
// 		}

// 		c.JSON(http.StatusCreated, gin.H{"id": id})
// 		// Se der erro de marshalling no JSON?
// 	}
// }

// func updateUnitOfMeasure(service unitofmeasure.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id := c.Param("id")

// 		if id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
// 			return
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 			return
// 		}

// 		var input presenter.UpdateUnitOfMeasure
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := service.UpdateUnitOfMeasure(&entity.UnitOfMeasure{
// 			ID:   idUUID,
// 			Name: input.Name,
// 		}); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"status": "unit of measure udpated"})
// 	}
// }

// func deleteUnitOfMeasure(service unitofmeasure.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id := c.Param("id")
// 		if id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
// 			return
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
// 			return
// 		}

// 		if err := service.DeleteUnitOfMeasure(idUUID); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting unit of measure"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"status": "unit of measure deleted"})
// 	}
// }
