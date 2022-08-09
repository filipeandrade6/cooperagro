package main

import (
	"log"
	"net/http"

	handler "github.com/filipeandrade6/cooperagro/cmd/api/handler/echo"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"

	//"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func main() {
	dataSourceName := "postgresql://postgres:postgres@localhost:5432/cooperagro"
	db, err := postgres.NewPostgresRepo(dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	baseProductService := baseproduct.NewService(db)
	// userService := user.NewService(db)
	// inventoryService := inventory.NewService(db)
	// productService := product.NewService(db)
	// unitOfMeasureService := unitofmeasure.NewService(db)

	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	handler.MakeBaseProductHandlers(e, baseProductService)

	// gin.MakeBaseProductHandlers(r, baseProductService)
	// gin.MakeUserHandlers(r, userService)
	// gin.MakeInventoryHandlers(r, inventoryService)
	// gin.MakeProductHandlers(r, productService)
	// gin.MakeUnitOfMeasureHandlers(r, unitOfMeasureService)

	// r.Run()

	e.Logger.Fatal(e.Start(":8080"))
}
