package main

import (
	"log"
	"net/http"

	"github.com/filipeandrade6/cooperagro/cmd/api/handler"
	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres"

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
	productService := product.NewService(baseProductService, db)
	unitOfMeasureService := unitofmeasure.NewService(db)
	userService := user.NewService(db)
	inventoryService := inventory.NewService(
		productService,
		unitOfMeasureService,
		userService,
		db,
	)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.MakeAuthHandlers(e, userService)

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
	handler.MakeProductHandlers(v1, productService)
	handler.MakeUnitOfMeasureHandlers(v1, unitOfMeasureService)
	handler.MakeUserHandlers(v1, userService)
	handler.MakeInventoryHandlers(v1, inventoryService)

	e.Logger.Fatal(e.Start(":8080"))
}
