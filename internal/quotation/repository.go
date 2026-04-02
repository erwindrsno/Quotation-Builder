package quotation

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Save(c context.Context, item Quotation) error {
	_, err := r.DB.ExecContext(c, saveQuery, item.Subject, item.QuoteNumber, item.Validity, item.DeliveryTime, item.DeliveryDestination, item.TermsOfPayment, item.Notes, item.Discount, item.ClientId)
	if err != nil {
		return err
	}
	return nil
}

// TODO: So far, only can query subject, need to add quote number in the future
func (r *Repository) FindPaginated(c context.Context, subject string, limit, offset int) ([]Quotation, error) {
	rows, err := r.DB.QueryContext(c, findPaginatedQuery, "%"+subject+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Quotation, limit)

	for rows.Next() {
		var item Quotation
		err := rows.Scan(&item.Id, &item.Subject, &item.QuoteNumber, &item.Validity, &item.DeliveryTime, &item.Deadline, &item.ClientId, &item.CreatedAt)
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

func (r *Repository) FindById(c context.Context, id uuid.UUID) (Quotation, error) {
	var item Quotation

	err := r.DB.QueryRowContext(c, findByIdQuery, id).Scan(
		&item.Id,
		&item.Subject,
		&item.QuoteNumber,
		&item.Validity,
		&item.DeliveryTime,
		&item.Deadline,
		&item.DeliveryDestination,
		&item.TermsOfPayment,
		&item.Notes,
		&item.Discount,
		&item.ClientId,
		&item.CreatedAt,
	)

	if err != nil {
		// Return an empty Quotation struct and the error
		return Quotation{}, err
	}

	return item, nil
}
