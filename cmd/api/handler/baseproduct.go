package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entities"
	"github.com/filipeandrade6/cooperagro/domain/usecases/baseproduct"
	"github.com/gin-gonic/gin"
)

func MakeBaseProductHandlers(r *gin.Engine, service baseproduct.UseCase) {
	r.GET("/v1/baseproduct", listBaseProduct(service))
}

// func searchBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		errorMessage := "error reading baseproduct"

// 		name := c.Param("name")
// 		data, err := service.SearchBaseProduct(name)
// 		if err != nil && !errors.Is(err, entities.ErrNotFound) {
// 			c.JSON(http.StatusInternalServerError, errorMessage)
// 			return
// 		}

// 		if data == nil {
// 			c.JSON(http.StatusNotFound, "not found")
// 			return
// 		}

// 		var toJ []*presenter.BaseProduct
// 		for _, d := range data {
// 			toJ = append(toJ, &presenter.BaseProduct{
// 				ID:   d.ID,
// 				Name: d.Name,
// 			})
// 		}
// 		c.JSON(http.StatusOK, toJ)

// 		// Se der erro de marshalling no JSON?
// 	}
// }

func listBaseProduct(service baseproduct.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading baseproduct"

		var data []*entities.BaseProduct
		var err error

		name := c.Query("name")
		switch {
		case name == "":
			data, err = service.ListBaseProduct()
		default:
			data, err = service.SearchBaseProduct(name)
		}

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, errorMessage)
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, nil)
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
