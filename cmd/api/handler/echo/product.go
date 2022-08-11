package echo

// import (
// 	"errors"
// 	"net/http"

// 	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware/echo"
// 	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
// 	"github.com/labstack/echo/v4"
// )

// func MakeProductHandlers(e *echo.Group, service product.UseCase) {
// 	e.POST("/products", createProduct(service), mid.AdminRequired)
// 	e.GET("/products", readProduct(service))
// 	e.GET("/products/:id", getProduct(service))
// 	e.PUT("/products/:id", updateProduct(service), mid.AdminRequired)
// 	e.DELETE("/products/:id", deleteProduct(service), mid.AdminRequired)
// }

// func createProduct(service product.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.EchoProduct
// 		if err := c.Bind(&input); err != nil {
// 			return echo.ErrBadRequest
// 		}

// 		idUIID, err := entity.StringToID(input.BaseProductID)
// 		if err != nil {
// 			return echo.ErrBadRequest
// 		}

// 		id, err := service.CreateProduct(input.Name, idUIID)
// 		if errors.Is(entity.ErrEntityAlreadyExists, err) {
// 			return c.NoContent(http.StatusConflict)
// 		}
// 		if errors.Is(entity.ErrInvalidEntity, err) {
// 			return echo.ErrBadRequest
// 		}
// 		if err != nil {
// 			return echo.ErrInternalServerError
// 		}

// 		return c.JSON(
// 			http.StatusCreated,
// 			echo.Map{"id": id.String()},
// 		)
// 	}
// }

// func getProduct(service product.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idUUID, err := entity.StringToID(c.Param("id"))
// 		if err != nil {
// 			return echo.ErrBadRequest
// 		}

// 		data, err := service.GetProductByID(idUUID)
// 		if errors.Is(err, entity.ErrNotFound) {
// 			return echo.ErrNotFound
// 		}
// 		if err != nil {
// 			return echo.ErrInternalServerError
// 		}

// 		return c.JSON(http.StatusOK, &presenter.Product{
// 			ID:   data.ID,
// 			Name: data.Name,
// 		})
// 	}
// }

// func readProduct(service product.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var data []*entity.Product
// 		var err error

// 		name := c.QueryParam("name")
// 		if name != "" {
// 			data, err = service.SearchProduct(name)
// 		} else {
// 			data, err = service.ListProduct()
// 		}

// 		if errors.Is(err, entity.ErrNotFound) {
// 			return echo.ErrNotFound
// 		}
// 		if err != nil {
// 			return echo.ErrInternalServerError
// 		}

// 		var out []*presenter.Product
// 		for _, d := range data {
// 			out = append(out, &presenter.Product{
// 				ID:   d.ID,
// 				Name: d.Name,
// 			})
// 		}

// 		return c.JSON(http.StatusOK, out)
// 	}
// }

// func updateProduct(service product.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idUUID, err := entity.StringToID(c.Param("id"))
// 		if err != nil {
// 			return echo.ErrBadRequest
// 		}

// 		var input presenter.Product
// 		if err := c.Bind(&input); err != nil {
// 			return echo.ErrInternalServerError
// 		}

// 		err = service.UpdateProduct(&entity.Product{
// 			ID:   idUUID,
// 			Name: input.Name,
// 		})
// 		switch {
// 		case errors.Is(entity.ErrInvalidEntity, err):
// 			return echo.ErrBadRequest

// 		case errors.Is(entity.ErrNotFound, err):
// 			return echo.ErrNotFound

// 		case errors.Is(entity.ErrEntityAlreadyExists, err):
// 			return c.NoContent(http.StatusConflict)

// 		case err != nil:
// 			return echo.ErrInternalServerError
// 		}

// 		return c.NoContent(http.StatusOK)
// 	}
// }

// func deleteProduct(service product.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idUUID, err := entity.StringToID(c.Param("id"))
// 		if err != nil {
// 			return echo.ErrBadRequest
// 		}

// 		err = service.DeleteProduct(idUUID)
// 		if errors.Is(err, entity.ErrNotFound) {
// 			return echo.ErrNotFound
// 		}
// 		if err != nil {
// 			return echo.ErrInternalServerError
// 		}

// 		return c.NoContent(http.StatusOK)
// 	}
// }
