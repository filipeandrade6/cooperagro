package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entities"
	"github.com/filipeandrade6/cooperagro/domain/usecases/customer"
	"github.com/gin-gonic/gin"
)

func MakeCustomerHandlers(r *gin.Engine, service customer.UseCase) {
	r.GET("/customer/:id", getCustomerByID(service))
	r.GET("/customer", listCustomer(service))
	r.POST("/customer", createCustomer(service))
	r.PUT("/customer/:id", updateCustomer(service))
	r.DELETE("/customer/:id", deleteCustomer(service))
}

func getCustomerByID(service customer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading customer"

		id, err := entities.StringToID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		data, err := service.GetCustomerByID(id)

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		c.JSON(http.StatusOK, &presenter.Customer{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Address:   data.Address,
			Phone:     data.Phone,
			Email:     data.Email,
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		})

		// Se der erro de marshalling no JSON?
	}
}

func listCustomer(service customer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading customer"

		var data []*entities.Customer
		var err error

		firstName := c.Query("first_name")
		switch {
		case firstName == "":
			data, err = service.ListCustomer()
		default:
			data, err = service.SearchCustomer(firstName)
		}

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		var toJ []*presenter.Customer
		for _, d := range data {
			toJ = append(toJ, &presenter.Customer{
				ID:        d.ID,
				FirstName: d.FirstName,
				LastName:  d.LastName,
				Address:   d.Address,
				Phone:     d.Phone,
				Email:     d.Email,
				Latitude:  d.Latitude,
				Longitude: d.Longitude,
			})
		}
		c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createCustomer(service customer.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input presenter.CreateCustomer
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
			return
		}

		id, err := service.CreateCustomer(
			input.FirstName,
			input.LastName,
			input.Address,
			input.Phone,
			input.Email,
			input.Latitude,
			input.Longitude,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "creating customer"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
		// Se der erro de marshalling no JSON?
	}
}

func updateCustomer(service customer.UseCase) gin.HandlerFunc {
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

		var input presenter.UpdateCustomer
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.UpdateCustomer(&entities.Customer{
			ID:        idUUID,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Address:   input.Address,
			Phone:     input.Phone,
			Email:     input.Email,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "customer udpated"})
	}
}

func deleteCustomer(service customer.UseCase) gin.HandlerFunc {
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

		if err := service.DeleteCustomer(idUUID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting customer"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "customer deleted"})
	}
}
