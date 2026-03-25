package user

const saveQuery = `
	INSERT INTO users
		(username, password, name, role_id, created_at) 
	VALUES 
		($1, $2, $3, $4, CURRENT_TIMESTAMP)
`

const findQuery = `
  SELECT id, username, name, role_id, created_at, updated_at
  FROM users 
  WHERE username ILIKE $1 OR name ILIKE $2
  ORDER BY created_at ASC, id ASC
  LIMIT $3 OFFSET $4;
`
