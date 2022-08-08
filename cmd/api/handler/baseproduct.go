package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/gin-gonic/gin"
)

func MakeBaseProductHandlers(r *gin.Engine, service baseproduct.UseCase) {
	r.GET("/baseproduct/:id", getBaseProductByID(service))
	r.GET("/baseproduct", listBaseProduct(service))
	r.POST("/baseproduct", createBaseProduct(service))
	r.PUT("/baseproduct/:id", updateBaseProduct(service))
	r.DELETE("/baseproduct/:id", deleteBaseProduct(service))
}

func getBaseProductByID(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading baseproduct"

		id, err := entity.StringToID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		data, err := service.GetBaseProductByID(id)

		if err != nil && !errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		c.JSON(http.StatusOK, &presenter.BaseProduct{
			ID:   data.ID,
			Name: data.Name,
		})

		// Se der erro de marshalling no JSON?
	}
}

func listBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading baseproduct"

		var data []*entity.BaseProduct
		var err error

		name := c.Query("name")
		switch {
		case name == "":
			data, err = service.ListBaseProduct()
		default:
			data, err = service.SearchBaseProduct(name)
		}

		if err != nil && !errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		var toJ []*presenter.BaseProduct
		for _, d := range data {
			toJ = append(toJ, &presenter.BaseProduct{
				ID:   d.ID,
				Name: d.Name,
			})
		}
		c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input presenter.CreateBaseProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
			return
		}

		id, err := service.CreateBaseProduct(input.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "creating base product"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
		// Se der erro de marshalling no JSON?
	}
}

func updateBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
			return
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var input presenter.UpdateBaseProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.UpdateBaseProduct(&entity.BaseProduct{
			ID:   idUUID,
			Name: input.Name,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "base product udpated"})
	}
}

func deleteBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
			return
		}

		idUUID, err := entity.StringToID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if err := service.DeleteBaseProduct(idUUID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting base product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "base product deleted"})
	}
}
