package postgres

import (
	"database/sql"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entities"
)

func (r *Repo) Create(e *entities.Product) (entities.ID, error) {
	stmt, err := r.db.Prepare(`
	INSERT INTO products (id, name, base_product_id, created_at) values (?, ?, ?, ?, )`)
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(
		e.ID,
		e.Name,
		e.BaseProduct,
		e.CreatedAt.Format(time.RFC3339), // TODO já nem tem o create na entidade
	)
	if err != nil {
		return e.ID, err
	}
	if err = stmt.Close(); err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *ProductPG) GetBase(id entities.ID) (*entities.Product, error) {
	stmt, err := r.db.Prepare(`
	SELECT id, name, base_product_id, created_at, updated_at FROM procuts WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	var p entities.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.BaseProduct, &p.CreatedAt, &p.UpdatedAt)
	}
	// TODO tratar esse erro

	return &p, nil
}

func (r *ProductPG)
