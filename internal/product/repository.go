package product

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Save(c context.Context, item Product) error {
	_, err := r.DB.ExecContext(c, saveQuery, item.Name, item.PartNumber, item.Description, item.BaseUnit, item.BasePrice, item.ManufacturerBrand)
	if err != nil {
		return err
	}
	return nil
}

// TODO: tobe fix by adding search by name and manufacturer brand
func (r *Repository) FindList(c context.Context, name string) ([]Product, error) {
	rows, err := r.DB.QueryContext(c, findListQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Product, 0)

	for rows.Next() {
		var item Product
		err := rows.Scan(&item.Id, &item.Name)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
