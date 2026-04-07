package client

const saveQuery = `
	INSERT INTO clients 
		(name, company_id, created_at) 
	VALUES 
		($1, $2, CURRENT_TIMESTAMP)
`

const findPaginatedQuery = `
	SELECT id, name, created_at
	FROM clients
  WHERE name ILIKE $1
  ORDER BY created_at ASC, id ASC
  LIMIT $2 OFFSET $3;
`

const findListQuery = `
	SELECT c.id, c.name AS client_name, co.name AS company_name
	FROM clients c
	INNER JOIN companies co
	ON c.company_id = co.id
  WHERE c.name ILIKE $1
	ORDER BY c.name ASC;
`
