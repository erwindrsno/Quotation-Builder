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

	itemStatuses := make([]ItemStatus, 0)

	for rows.Next() {
		var itemStatus ItemStatus
		err := rows.Scan(&itemStatus.Id, &itemStatus.Name)
		if err != nil {
			return nil, err
		}
		itemStatuses = append(itemStatuses, itemStatus)
	}

	// check for errors that happened during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return itemStatuses, nil
}
