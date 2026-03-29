package client

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Save(c context.Context, name string, companyId uuid.UUID) error {
	_, err := r.DB.ExecContext(c, saveQuery, name, companyId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindList(c context.Context, name string) ([]Client, error) {
	rows, err := r.DB.QueryContext(c, findListQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Client, 0)

	for rows.Next() {
		var item Client
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

func (r *Repository) FindPaginated(c context.Context, name string, limit, offset int) ([]Client, error) {
	rows, err := r.DB.QueryContext(c, findPaginatedQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Client, 0)

	for rows.Next() {
		var item Client
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
