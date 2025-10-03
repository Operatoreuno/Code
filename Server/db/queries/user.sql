-- name: CheckUserEmailExists :one
SELECT COUNT(*) FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (email, password, name, surname, stripe_customer_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;