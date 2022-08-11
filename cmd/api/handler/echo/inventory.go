package echo

// import (
// 	"errors"
// 	"net/http"

// 	mid "github.com/filipeandrade6/cooperagro/cmd/api/middleware/echo"
// 	"github.com/filipeandrade6/cooperagro/cmd/api/presenter"
// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory"
// 	"github.com/filipeandrade6/cooperagro/infra/auth"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// )

// func MakeInventoryHandlers(e *echo.Group, service inventory.UseCase) {
// 	e.POST("/inventories", createInventory(service), mid.AdminRequired, mid.NeedUserID)
// 	e.GET("/inventories", readInventory(service))
// 	e.GET("/inventories/:id", getInventory(service))
// 	e.PUT("/inventories/:id", updateInventory(service), mid.AdminRequired, mid.NeedUserID)
// 	e.DELETE("/inventories/:id", deleteInventory(service), mid.AdminRequired, mid.NeedUserID)
// }

// func createInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input presenter.Inventory
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		id, err := service.CreateInventory()

// 		if errors.Is(entity.ErrEntityAlreadyExists, err) {
// 			return c.JSON(
// 				http.StatusConflict,
// 				echo.Map{"status": "base product already exists"},
// 			)
// 		}
// 		if errors.Is(entity.ErrInvalidEntity, err) {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid parameters"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(
// 			http.StatusCreated,
// 			echo.Map{"id": id.String()},
// 		)
// 	}
// }

// func getInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id := c.Param("id")
// 		if id == "" {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "empty id"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid id"},
// 			)
// 		}

// 		data, err := service.GetInventoryByID(idUUID)
// 		if errors.Is(err, entity.ErrNotFound) {
// 			return c.JSON(
// 				http.StatusNotFound,
// 				echo.Map{"status": "base product not found"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()}, // TODO - não expor o erro ao usuŕio?
// 			)
// 		}

// 		return c.JSON(http.StatusOK, &presenter.Inventory{
// 			ID:   data.ID,
// 			Name: data.Name,
// 		})
// 	}
// }

// func readInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var data []*entity.Inventory
// 		var err error

// 		name := c.QueryParam("name")
// 		if name != "" {
// 			data, err = service.SearchInventory(name)
// 		} else {
// 			data, err = service.ListInventory()
// 		}

// 		if errors.Is(err, entity.ErrNotFound) {
// 			return c.JSON(
// 				http.StatusNotFound,
// 				echo.Map{"status": "base products not found"},
// 			)
// 		}
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		var out []*presenter.Inventory
// 		for _, d := range data {
// 			out = append(out, &presenter.Inventory{
// 				ID:   d.ID,
// 				Name: d.Name,
// 			})
// 		}

// 		return c.JSON(http.StatusOK, out)
// 	}
// }

// func updateInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(*auth.Claims)

// 		if !claims.Authorized("admin") {
// 			return echo.ErrForbidden
// 		}

// 		id := c.Param("id")

// 		if id == "" {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "empty id"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid id"},
// 			)
// 		}

// 		var input presenter.Inventory
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": "could not get values from the request"},
// 			)
// 		}

// 		if err := service.UpdateInventory(&entity.Inventory{
// 			ID:   idUUID,
// 			Name: input.Name,
// 		}); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"status": "base product udpated"})
// 	}
// }

// func deleteInventory(service inventory.UseCase) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(*auth.Claims)

// 		if !claims.Authorized("admin") {
// 			return echo.ErrForbidden
// 		}

// 		id := c.Param("id")

// 		if id == "" {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "empty id"},
// 			)
// 		}

// 		idUUID, err := entity.StringToID(id)
// 		if err != nil {
// 			return c.JSON(
// 				http.StatusBadRequest,
// 				echo.Map{"status": "invalid id"},
// 			)
// 		}

// 		if err := service.DeleteInventory(idUUID); err != nil {
// 			return c.JSON(
// 				http.StatusInternalServerError,
// 				echo.Map{"status": err.Error()},
// 			)
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"status": "base product deleted"})
// 	}
// }
