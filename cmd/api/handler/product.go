package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entities"
	"github.com/filipeandrade6/cooperagro/domain/usecases/product"
	"github.com/gin-gonic/gin"
)

func MakeProductHandlers(r *gin.Engine, service product.UseCase) {
	r.GET("/product/:id", getProductByID(service))
	r.GET("/product", listProduct(service))
	r.POST("/product", createProduct(service))
	r.PUT("/product/:id", updateProduct(service))
	r.DELETE("/product/:id", deleteProduct(service))
}

func getProductByID(service product.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading product"

		id, err := entities.StringToID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		data, err := service.GetProductByID(id)

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		c.JSON(http.StatusOK, &presenter.Product{
			ID:          data.ID,
			Name:        data.Name,
			BaseProduct: data.BaseProduct,
		})

		// Se der erro de marshalling no JSON?
	}
}

func listProduct(service product.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading product"

		data, err := service.ListProduct()

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		var toJ []*presenter.Product
		for _, d := range data {
			toJ = append(toJ, &presenter.Product{
				ID:          d.ID,
				Name:        d.Name,
				BaseProduct: d.BaseProduct,
			})
		}
		c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createProduct(service product.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input presenter.CreateProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
			return
		}

		id, err := service.CreateProduct(
			input.Name,
			input.BaseProduct,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "creating product"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
		// Se der erro de marshalling no JSON?
	}
}

func updateProduct(service product.UseCase) gin.HandlerFunc {
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

		var input presenter.UpdateProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.UpdateProduct(&entities.Product{
			ID:          idUUID,
			Name:        input.Name,
			BaseProduct: input.BaseProduct,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "product udpated"})
	}
}

func deleteProduct(service product.UseCase) gin.HandlerFunc {
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

		if err := service.DeleteProduct(idUUID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "product deleted"})
	}
}
