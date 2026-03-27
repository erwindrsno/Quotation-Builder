package item_status

const saveQuery = `
	INSERT INTO item_statuses 
		(name, created_at) 
	VALUES 
		($1, CURRENT_TIMESTAMP)
`

const findListQuery = `
	SELECT id, name, created_at
	FROM item_statuses
  WHERE name ILIKE $1
`
