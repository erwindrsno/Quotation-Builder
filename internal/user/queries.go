package user

const createQuery = `
	INSERT INTO users
		(username, password, name, role_id, created_at) 
	VALUES 
		($1, $2, $3, $4, CURRENT_TIMESTAMP)
`
