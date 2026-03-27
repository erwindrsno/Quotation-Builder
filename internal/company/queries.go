package company

const saveQuery = `
	INSERT INTO companies 
		(name, created_at) 
	VALUES 
		($1, CURRENT_TIMESTAMP)
`

const findPaginatedQuery = `
	SELECT id, name, created_at
	FROM companies
  WHERE name ILIKE $1
  ORDER BY created_at ASC, id ASC
  LIMIT $2 OFFSET $3;
`

const findListQuery = `
	SELECT id, name, created_at
	FROM companies
  WHERE name ILIKE $1
`
