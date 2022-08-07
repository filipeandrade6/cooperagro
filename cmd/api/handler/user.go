package handler

import (
	"errors"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
	"github.com/filipeandrade6/cooperagro/domain/entities"
	"github.com/filipeandrade6/cooperagro/domain/usecases/user"
	"github.com/gin-gonic/gin"
)

func MakeUserHandlers(r *gin.Engine, service user.UseCase) {
	r.GET("/user/:id", getUserByID(service))
	r.GET("/user", listUser(service))
	r.POST("/user", createUser(service))
	r.PUT("/user/:id", updateUser(service))
	r.DELETE("/user/:id", deleteUser(service))
}

func getUserByID(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading user"

		id, err := entities.StringToID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		data, err := service.GetUserByID(id)

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		c.JSON(http.StatusOK, &presenter.User{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Address:   data.Address,
			Phone:     data.Phone,
			Email:     data.Email,
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
			Role:      data.Role,
		})

		// Se der erro de marshalling no JSON?
	}
}

func listUser(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorMessage := "error reading user"

		var data []*entities.User
		var err error

		firstName := c.Query("first_name")
		switch {
		case firstName == "":
			data, err = service.ListUser()
		default:
			data, err = service.SearchUser(firstName)
		}

		if err != nil && !errors.Is(err, entities.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": errorMessage})
			return
		}

		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
			return
		}

		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:        d.ID,
				FirstName: d.FirstName,
				LastName:  d.LastName,
				Address:   d.Address,
				Phone:     d.Phone,
				Email:     d.Email,
				Latitude:  d.Latitude,
				Longitude: d.Longitude,
				Role:      d.Role,
			})
		}
		c.JSON(http.StatusOK, toJ)

		// Se der erro de marshalling no JSON?
	}
}

func createUser(service user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input presenter.CreateUser
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO aplicar para outros
			return
		}

		id, err := service.CreateUser(
			input.FirstName,
			input.LastName,
			input.Address,
			input.Phone,
			input.Email,
			input.Latitude,
			input.Longitude,
			input.Role,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "creating user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
		// Se der erro de marshalling no JSON?
	}
}

func updateUser(service user.UseCase) gin.HandlerFunc {
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

		var input presenter.UpdateUser
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.UpdateUser(&entities.User{
			ID:        idUUID,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Address:   input.Address,
			Phone:     input.Phone,
			Email:     input.Email,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Role:      input.Role,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "user udpated"})
	}
}

func deleteUser(service user.UseCase) gin.HandlerFunc {
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

		if err := service.DeleteUser(idUUID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "deleting user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
	}
}
