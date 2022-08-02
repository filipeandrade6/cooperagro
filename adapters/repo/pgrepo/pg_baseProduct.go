package pgrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/vingarcia/ksql"
)

func changeBaseProductName(ctx context.Context, db ksql.Provider, baseProductID int, newBaseProductName string) error {
	return db.Transaction(ctx, func(db ksql.Provider) error {
		baseProduct, err := getBaseProduct(ctx, db, baseProductID)
		if err != nil {
			return err
		}

		// If there is nothing to do, just return
		if baseProduct.Name == newBaseProductName {
			return nil
		}

		_, err = getBaseProductByName(ctx, db, newBaseProductName)
		if err != ksql.ErrRecordNotFound {
			return fmt.Errorf("can't change base product name to '%s': this name is already used by other base product", newBaseProductName)
		}
		if err != nil {
			return err
		}

		baseProduct.Name = newBaseProductName
		_, err = upsertBaseProduct(ctx, db, baseProduct)
		return err
	})
}

// Keeping the implementation deatached like this and passing the database provider interface
// as an argument allows you to include several diferent calls in a same transaction.
func upsertBaseProduct(ctx context.Context, db ksql.Provider, baseProduct domain.BaseProduct) (baseProductID int, _ error) {
	now := time.Now()
	baseProduct.UpdatedAt = &now
	err := db.Patch(ctx, domain.BaseProductsTable, &baseProduct)
	if err == ksql.ErrRecordNotFound {
		baseProduct.CreatedAt = &now
		err = db.Insert(ctx, domain.BaseProductsTable, &baseProduct)
	}
	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving base product", map[string]interface{}{
			"base_product": baseProduct,
			"error":        err.Error(),
		})
	}

	return baseProductID, nil
}

func getBaseProduct(ctx context.Context, db ksql.Provider, baseProductID int) (domain.BaseProduct, error) {
	var baseProduct domain.BaseProduct
	err := db.QueryOne(ctx, &baseProduct, "FROM base_products WHERE id = $1", baseProductID)
	if err == ksql.ErrRecordNotFound {
		return domain.BaseProduct{}, domain.NotFoundErr("no base product found with provided id", map[string]interface{}{
			"base_product_id": baseProductID,
		})
	}
	if err != nil {
		return domain.BaseProduct{}, domain.InternalErr("unexpected error when fetching base product", map[string]interface{}{
			"base_product_id": baseProductID,
			"error":           err.Error(),
		})
	}

	return baseProduct, nil
}

func getBaseProductByName(ctx context.Context, db ksql.Provider, name string) (domain.BaseProduct, error) {
	var baseProduct domain.BaseProduct
	err := db.QueryOne(ctx, &baseProduct, "FROM base_products WHERE name = $1", name)
	if err == ksql.ErrRecordNotFound {
		return domain.BaseProduct{}, domain.NotFoundErr("no base product found with provided name", map[string]interface{}{
			"name": name,
		})
	}
	if err != nil {
		return domain.BaseProduct{}, domain.InternalErr("unexpected error when fetching base product by name", map[string]interface{}{
			"name":  name,
			"error": err.Error(),
		})
	}

	return baseProduct, nil
}
