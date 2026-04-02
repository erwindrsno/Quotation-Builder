package quotation

const saveQuery = `
	INSERT INTO quotations 
		(subject, quote_number, validity, delivery_time, deadline, delivery_destination, terms_of_payment, notes, discount, client_id, created_at) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP)
`

const findPaginatedQuery = `
	SELECT id, subject, quote_number, validity, delivery_time, deadline, client_id, created_at)
	FROM quotations
  WHERE subject ILIKE $1
  ORDER BY created_at ASC, id ASC
  LIMIT $2 OFFSET $3;
`

const findByIdQuery = `
	SELECT 
		subject, quote_number, validity, delivery_time, delivery_destination, terms_of_payment, notes, discount, client_id, created_at
	FROM quotations
	WHERE id = $1
`
