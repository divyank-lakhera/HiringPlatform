-- name: CreateUser :one
INSERT INTO users (
    username,
    password
) values (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1;

-- name: UpdateCredentials :exec
UPDATE users
SET password = $2
WHERE id = $1;