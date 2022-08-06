package productsctrl

import (
	"encoding/json"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/filipeandrade6/cooperagro/domain/products"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	productsService products.Service
}

func NewController(productsService products.Service) Controller {
	return Controller{
		productsService: productsService,
	}
}

func (c Controller) UpsertProduct(ctx *fiber.Ctx) error {
	// Intermediary structure so I don't expose my internal
	// product representation to the outside world
	var product struct {
		ProductID int    `json:"product_id"`
		Name      string `json:"name"`
	}
	err := json.Unmarshal(ctx.Body(), &product)
	if err != nil {
		return domain.BadRequestErr("unable to parse payload as JSON", map[string]interface{}{
			"payload": string(ctx.Body()),
			"error":   err.Error(),
		})
	}

	productID, err := c.productsService.UpsertProduct(ctx.Context(), domain.Product{
		// Showcasing that my internal model might differ from the API,
		// in this case the internal name for the ID attribute is just ID not `ProductID`
		ID:   product.ProductID,
		Name: product.Name,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"status":     "success",
		"product_id": productID,
	})
}

func (c Controller) GetProduct(ctx *fiber.Ctx) error {
	productID, err := ctx.ParamsInt("id")
	if err != nil {
		return domain.BadRequestErr("the input product id is not a valid integer", map[string]interface{}{
			"received_id": ctx.Params(":id"),
		})
	}

	product, err := c.productsService.GetProduct(ctx.Context(), productID)
	if err != nil {
		return err
	}

	// Again using intermediary structs (or a map) is useful for decoupling
	// the internal entities from what is exposed on the web:
	return ctx.JSON(map[string]interface{}{
		"id":   productID,
		"name": product.Name,
	})
}
