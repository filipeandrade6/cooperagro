package main

import (
	"log"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/handler"
	"github.com/filipeandrade6/cooperagro/domain/usecases/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecases/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecases/product"
	"github.com/filipeandrade6/cooperagro/domain/usecases/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecases/user"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	dataSourceName := "postgresql://postgres:postgres@localhost:5432/cooperagro"
	db, err := postgres.NewPostgresRepo(dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	baseProductService := baseproduct.NewService(db)
	userService := user.NewService(db)
	inventoryService := inventory.NewService(db)
	productService := product.NewService(db)
	unitOfMeasureService := unitofmeasure.NewService(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handler.MakeBaseProductHandlers(r, baseProductService)
	handler.MakeUserHandlers(r, userService)
	handler.MakeInventoryHandlers(r, inventoryService)
	handler.MakeProductHandlers(r, productService)
	handler.MakeUnitOfMeasureHandlers(r, unitOfMeasureService)

	r.Run()
}
