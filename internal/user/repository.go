package user

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Create(c context.Context, u *Register) error {
	_, err := r.DB.ExecContext(c, createQuery, u.Username, u.Password, u.Name, u.RoleId)
	if err != nil {
		return err
	}
	return nil
}
