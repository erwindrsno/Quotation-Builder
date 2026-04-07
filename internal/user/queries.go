package user

const saveQuery = `
	INSERT INTO users
		(username, password, name, role_id, created_at)
	VALUES
		($1, $2, $3, $4, CURRENT_TIMESTAMP)
`

const findQuery = `
  SELECT u.id, u.username, u.name, u.role_id, u.created_at, u.updated_at, r.name AS role_name
  FROM users u
  INNER JOIN roles r ON u.role_id = r.id
  WHERE u.username ILIKE $1 OR u.name ILIKE $2
  ORDER BY u.created_at ASC, u.id ASC
  LIMIT $3 OFFSET $4;
`

const findPasswordByUsername = `
  SELECT password
  FROM users
  WHERE username = $1
`
