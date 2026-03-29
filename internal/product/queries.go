package product

const saveQuery = `
	INSERT INTO products 
		(name, part_number, description, base_units, manufacturer_brand, created_at) 
	VALUES 
		($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
`

const findPaginatedQuery = `
	SELECT id, name, part_number, description, base_units, manufacturer_brand, created_at)
	FROM products
  WHERE name ILIKE $1
  ORDER BY created_at ASC, id ASC
  LIMIT $2 OFFSET $3;
`

const findListQuery = `
	SELECT id, name
	FROM companies
  WHERE name ILIKE $1
	ORDER BY name ASC;
`
