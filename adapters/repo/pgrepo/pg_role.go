package pgrepo

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/cooperagro/domain"
	"github.com/vingarcia/ksql"
)

func changeRoleName(ctx context.Context, db ksql.Provider, roleID int, newRoleName string) error {
	return db.Transaction(ctx, func(db ksql.Provider) error {
		role, err := getRoleByID(ctx, db, roleID)
		if err != nil {
			return err
		}

		// If there is nothing to do, just return
		if role.Name == newRoleName {
			return nil
		}

		_, err = getRoleByName(ctx, db, newRoleName)
		if err != ksql.ErrRecordNotFound {
			return fmt.Errorf("can't change role name to '%s': this name is already used by other role", newRoleName)
		}
		if err != nil {
			return err
		}

		role.Name = newRoleName
		_, err = upsertRole(ctx, db, role)
		return err
	})
}

// Keeping the implementation deatached like this and passing the database provider interface
// as an argument allows you to include several diferent calls in a same transaction.
func upsertRole(ctx context.Context, db ksql.Provider, role domain.Role) (roleID int, _ error) {
	err := db.Patch(ctx, domain.RolesTable, &role)
	if err == ksql.ErrRecordNotFound {
		err = db.Insert(ctx, domain.RolesTable, &role)
	}
	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving role", map[string]interface{}{
			"role":  role,
			"error": err.Error(),
		})
	}

	return roleID, nil
}

func getRoleByID(ctx context.Context, db ksql.Provider, roleID int) (domain.Role, error) {
	var role domain.Role
	err := db.QueryOne(ctx, &role, "FROM roles WHERE id = $1", roleID)
	if err == ksql.ErrRecordNotFound {
		return domain.Role{}, domain.NotFoundErr("no role found with provided id", map[string]interface{}{
			"role_id": roleID,
		})
	}
	if err != nil {
		return domain.Role{}, domain.InternalErr("unexpected error when fetching role", map[string]interface{}{
			"role_id": roleID,
			"error":   err.Error(),
		})
	}

	return role, nil
}

func getRoleByName(ctx context.Context, db ksql.Provider, name string) (domain.Role, error) {
	var role domain.Role
	err := db.QueryOne(ctx, &role, "FROM roles WHERE name = $1", name)
	if err == ksql.ErrRecordNotFound {
		return domain.Role{}, domain.NotFoundErr("no role found with provided name", map[string]interface{}{
			"name": name,
		})
	}
	if err != nil {
		return domain.Role{}, domain.InternalErr("unexpected error when fetching role by name", map[string]interface{}{
			"name":  name,
			"error": err.Error(),
		})
	}

	return role, nil
}
