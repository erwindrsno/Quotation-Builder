package user

import (
	"context"
	"database/sql"
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

	users := make([]User, 0, limit)

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Name, &u.RoleId, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) FindStoredHashByUsername(c context.Context, username string) (string, error) {
	row := r.DB.QueryRowContext(c, findPasswordByUsername, username)

	var password string

	if err := row.Scan(&password); err == sql.ErrNoRows {
		return "", err
	}
	return password, nil
}
