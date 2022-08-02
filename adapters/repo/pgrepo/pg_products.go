package pgrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/vingarcia/ksql"
)

func changeProductName(ctx context.Context, db ksql.Provider, productID int, newProductName string) error {
	return db.Transaction(ctx, func(db ksql.Provider) error {
		product, err := getProduct(ctx, db, productID)
		if err != nil {
			return err
		}

		// If there is nothing to do, just return
		if product.Name == newProductName {
			return nil
		}

		_, err = getProductByName(ctx, db, newProductName)
		if err != ksql.ErrRecordNotFound {
			return fmt.Errorf("can't change product name to '%s': this name is already used by other product", newProductName)
		}
		if err != nil {
			return err
		}

		product.Name = newProductName
		_, err = upsertProduct(ctx, db, product)
		return err
	})
}

// Keeping the implementation deatached like this and passing the database provider interface
// as an argument allows you to include several diferent calls in a same transaction.
func upsertProduct(ctx context.Context, db ksql.Provider, product domain.Product) (productID int, _ error) {
	now := time.Now()
	product.UpdatedAt = &now
	err := db.Patch(ctx, domain.ProductsTable, &product)
	if err == ksql.ErrRecordNotFound {
		product.CreatedAt = &now
		err = db.Insert(ctx, domain.ProductsTable, &product)
	}
	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving product", map[string]interface{}{
			"product": product,
			"error":   err.Error(),
		})
	}

	return productID, nil
}

func getProduct(ctx context.Context, db ksql.Provider, productID int) (domain.Product, error) {
	var product domain.Product
	err := db.QueryOne(ctx, &product, "FROM products WHERE id = $1", productID)
	if err == ksql.ErrRecordNotFound {
		return domain.Product{}, domain.NotFoundErr("no product found with provided id", map[string]interface{}{
			"product_id": productID,
		})
	}
	if err != nil {
		return domain.Product{}, domain.InternalErr("unexpected error when fetching product", map[string]interface{}{
			"product_id": productID,
			"error":      err.Error(),
		})
	}

	return product, nil
}

func getProductByName(ctx context.Context, db ksql.Provider, name string) (domain.Product, error) {
	var product domain.Product
	err := db.QueryOne(ctx, &product, "FROM products WHERE name = $1", name)
	if err == ksql.ErrRecordNotFound {
		return domain.Product{}, domain.NotFoundErr("no product found with provided name", map[string]interface{}{
			"name": name,
		})
	}
	if err != nil {
		return domain.Product{}, domain.InternalErr("unexpected error when fetching product by name", map[string]interface{}{
			"name":  name,
			"error": err.Error(),
		})
	}

	return product, nil
}
