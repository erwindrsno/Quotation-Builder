package item_status

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Save(c context.Context, name string) error {
	_, err := r.DB.ExecContext(c, saveQuery, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindList(c context.Context, name string) ([]ItemStatus, error) {
	rows, err := r.DB.QueryContext(c, findListQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]ItemStatus, 0)

	for rows.Next() {
		var item ItemStatus
		err := rows.Scan(&item.Id, &item.Name, &item.CreatedAt)
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
