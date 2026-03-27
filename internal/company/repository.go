package company

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

func (r *Repository) FindList(c context.Context, name string) ([]Company, error) {
	rows, err := r.DB.QueryContext(c, findListQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]Company, 0)

	for rows.Next() {
		var company Company
		err := rows.Scan(&company.Id, &company.Name, &company.CreatedAt)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *Repository) FindPaginated(c context.Context, name string, limit, offset int) ([]Company, error) {
	rows, err := r.DB.QueryContext(c, findPaginatedQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]Company, 0)

	for rows.Next() {
		var company Company
		err := rows.Scan(&company.Id, &company.Name, &company.CreatedAt)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
