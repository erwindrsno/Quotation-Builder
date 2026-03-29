package role

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

func (r *Repository) Find(c context.Context) ([]Role, error) {
	rows, err := r.DB.QueryContext(c, findQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Role, 0)

	for rows.Next() {
		var item Role
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
