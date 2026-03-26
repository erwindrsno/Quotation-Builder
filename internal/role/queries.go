package role

const saveQuery = `
	INSERT INTO roles
		(name,created_at) 
	VALUES 
		($1, CURRENT_TIMESTAMP)
`

const findQuery = `
	SELECT id, name, created_at
	FROM roles
`
