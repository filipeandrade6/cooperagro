package v1

import (
	"github.com/filipeandrade6/cooperagro/cmd/api/v1/handler"
	mid "github.com/filipeandrade6/cooperagro/cmd/api/v1/middleware"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
	"github.com/filipeandrade6/cooperagro/infra/auth"
	"github.com/filipeandrade6/cooperagro/infra/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterHandlers(e *echo.Echo, db repository.Repository) {
	// =================================================================================
	// services

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

	// login endpoint without JWT middleware
	handler.MakeAuthHandlers(e, userService)

	// =================================================================================
	// v1

	v1 := e.Group("/v1")
	v1.Use(
		middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &auth.Claims{},
			SigningKey: []byte("secret"),
		}),
		mid.ClaimsContext,
	)

	handler.MakeBaseProductHandlers(v1, baseProductService)
	handler.MakeProductHandlers(v1, productService)
	handler.MakeUnitOfMeasureHandlers(v1, unitOfMeasureService)
	handler.MakeUserHandlers(v1, userService)
	handler.MakeInventoryHandlers(v1, inventoryService)
}
