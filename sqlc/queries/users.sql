-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetTotalUsersCount :one
SELECT COUNT(*) FROM users;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  email,name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;