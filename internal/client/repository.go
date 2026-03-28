package client

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

func (r *Repository) FindList(c context.Context, name string) ([]Client, error) {
	rows, err := r.DB.QueryContext(c, findListQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := make([]Client, 0)

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.Id, &client.Name)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *Repository) FindPaginated(c context.Context, name string, limit, offset int) ([]Client, error) {
	rows, err := r.DB.QueryContext(c, findPaginatedQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := make([]Client, 0)

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.Id, &client.Name, &client.CreatedAt)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
