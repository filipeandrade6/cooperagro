package main

import (
	"log"
	"net/http"

	handler "github.com/filipeandrade6/cooperagro/cmd/api/handler/echo"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"

	//"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dataSourceName := "postgresql://postgres:postgres@localhost:5432/cooperagro"
	db, err := postgres.NewPostgresRepo(dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	baseProductService := baseproduct.NewService(db)
	userService := user.NewService(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := middleware.JWTConfig{
		Claims:     &auth.Claims{},
		SigningKey: []byte("secret"),
	}
	e.Use(middleware.JWTWithConfig(config))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/api/v1")

	handler.MakeAuthHandlers(v1, userService)
	handler.MakeBaseProductHandlers(v1, baseProductService)

	// gin.MakeBaseProductHandlers(r, baseProductService)
	// gin.MakeUserHandlers(r, userService)
	// gin.MakeInventoryHandlers(r, inventoryService)
	// gin.MakeProductHandlers(r, productService)
	// gin.MakeUnitOfMeasureHandlers(r, unitOfMeasureService)

	// r.Run()

	e.Logger.Fatal(e.Start(":8080"))
}
