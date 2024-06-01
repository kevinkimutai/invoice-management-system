-- name: GetCompany :one
SELECT * FROM company
WHERE company_id = $1 LIMIT 1;

-- name: ListCompany :many
SELECT * FROM company
ORDER BY company_name
LIMIT $1 OFFSET $2;

-- name: GetTotalCompaniesCount :one
SELECT COUNT(*) FROM company;

-- name: CreateCompany :one
INSERT INTO company (
  logo_url, user_id, company_name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateCompany :exec
UPDATE company
  set logo_url = $2,
  company_name = $3
WHERE company_id = $1;

-- name: DeleteCompany :exec
DELETE FROM company
WHERE company_id = $1;