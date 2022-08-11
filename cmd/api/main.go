package main

import (
	"log"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/handler"
	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
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
	unitOfMeasureService := unitofmeasure.NewService(db)
	userService := user.NewService(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.MakeAuthHandlers(e, userService)

	// e.Use(middleware.JWTWithConfig(config))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/api/v1")

	config := middleware.JWTConfig{
		Claims:     &auth.Claims{},
		SigningKey: []byte("secret"),
	}
	v1.Use(middleware.JWTWithConfig(config))
	v1.Use(mid.ClaimsContext)

	handler.MakeBaseProductHandlers(v1, baseProductService)

	// gin.MakeBaseProductHandlers(r, baseProductService)
	handler.MakeUserHandlers(v1, userService)
	// gin.MakeInventoryHandlers(r, inventoryService)
	// gin.MakeProductHandlers(r, productService)
	handler.MakeUnitOfMeasureHandlers(v1, unitOfMeasureService)

	// r.Run()

	e.Logger.Fatal(e.Start(":8080"))
}
