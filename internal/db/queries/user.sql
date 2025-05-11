-- name: CreateUser :one
INSERT INTO users (id, email, password, phone, full_name, created_at)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, NOW())
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;