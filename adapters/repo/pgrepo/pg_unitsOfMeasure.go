package pgrepo

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/vingarcia/ksql"
)

func upsertUnitOfMeasure(ctx context.Context, db ksql.Provider, unitOfMeasure domain.UnitOfMeasure) (unitOfMeasureID int, _ error) {
	err := db.Patch(ctx, domain.UnitsOfMeasureTable, &unitOfMeasure)
	if err == ksql.ErrRecordNotFound {
		err = db.Insert(ctx, domain.UnitsOfMeasureTable, &unitOfMeasure)
	}
	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving unit of measure", map[string]interface{}{
			"unit_of_measure": unitOfMeasure,
			"error":           err.Error(),
		})
	}

	return unitOfMeasureID, nil
}

func changeUnitOfMeasureName(ctx context.Context, db ksql.Provider, unitOfMeasureID int, newUnitOfMeasureName string) error {
	return db.Transaction(ctx, func(db ksql.Provider) error {
		unitOfMeasure, err := getUnitOfMeasureByID(ctx, db, unitOfMeasureID)
		if err != nil {
			return err
		}

		if unitOfMeasure.Name == newUnitOfMeasureName {
			return nil
		}

		_, err = getUnitOfMeasureByName(ctx, db, newUnitOfMeasureName)
		if err != ksql.ErrRecordNotFound {
			return fmt.Errorf("can't change unit of measure name to '%s': this name is already used by other unit of measure", newUnitOfMeasureName)
		}
		if err != nil {
			return err
		}

		unitOfMeasure.Name = newUnitOfMeasureName
		_, err = upsertUnitOfMeasure(ctx, db, unitOfMeasure)
		return err
	})
}

func getUnitOfMeasureByID(ctx context.Context, db ksql.Provider, unitOfMeasureID int) (domain.UnitOfMeasure, error) {
	var unitOfMeasure domain.UnitOfMeasure
	err := db.QueryOne(ctx, &unitOfMeasure, "FROM units_of_measure WHERE id = $1", unitOfMeasureID)
	if err == ksql.ErrRecordNotFound {
		return domain.UnitOfMeasure{}, domain.NotFoundErr("no unit of measure found with provided id", map[string]interface{}{
			"unit_of_measure": unitOfMeasureID,
		})
	}
	if err != nil {
		return domain.UnitOfMeasure{}, domain.InternalErr("unexpected error when fetching unit of measure", map[string]interface{}{
			"unit_of_measure_id": unitOfMeasureID,
			"error":              err.Error(),
		})
	}

	return unitOfMeasure, nil
}

func getUnitOfMeasureByName(ctx context.Context, db ksql.Provider, name string) (domain.UnitOfMeasure, error) {
	var unitOfMeasure domain.UnitOfMeasure
	err := db.QueryOne(ctx, &unitOfMeasure, "FROM unit_of_measure WHERE name = $1", name)
	if err == ksql.ErrRecordNotFound {
		return domain.UnitOfMeasure{}, domain.NotFoundErr("no unit of measure found with provided name", map[string]interface{}{
			"name": name,
		})
	}
	if err != nil {
		return domain.UnitOfMeasure{}, domain.InternalErr("unexpected error when fetching unit of measure by name", map[string]interface{}{
			"name":  name,
			"error": err.Error(),
		})
	}

	return unitOfMeasure, nil
}
