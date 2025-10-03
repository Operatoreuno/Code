-- name: GetAdminByID :one
SELECT * FROM admins WHERE id = $1;

-- name: GetAdminByEmail :one
SELECT * FROM admins WHERE email = $1;