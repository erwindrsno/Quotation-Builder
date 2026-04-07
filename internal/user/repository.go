package user

import (
	"context"
	"database/sql"

	"github.com/erwindrsno/Quotation-Builder/internal/role"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Save(c context.Context, u *Register) error {
	_, err := r.DB.ExecContext(c, saveQuery, u.Username, u.Password, u.Name, u.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Find(c context.Context, name string, limit, offset int) ([]User, error) {
	rows, err := r.DB.QueryContext(c, findQuery, "%"+name+"%", "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]User, 0, limit)

	for rows.Next() {
		var item User
		item.Role = &role.Role{}
		err := rows.Scan(&item.Id, &item.Username, &item.Name, &item.RoleId, &item.CreatedAt, &item.UpdatedAt, &item.Role.Name)
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

func (r *Repository) FindStoredHashByUsername(c context.Context, username string) (string, error) {
	row := r.DB.QueryRowContext(c, findPasswordByUsername, username)

	var password string

	err := row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
