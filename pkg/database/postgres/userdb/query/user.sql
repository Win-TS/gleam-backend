-- name: CreateUser :one
INSERT INTO users (
    username,
    firstname,
    lastname,
    phone_no,
    email,
    nationality,
    age,
    birthday,
    gender,
    photourl
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1 
FOR NO KEY UPDATE;

-- name: GetLatestId :one
SELECT COALESCE(MAX(id), 0)::integer FROM users;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ChangeUsername :exec
UPDATE users SET username = $2
WHERE id = $1;

-- name: ChangePhoneNo :exec
UPDATE users SET phone_no = $2
WHERE id = $1;

-- name: UpdateProfile :exec
UPDATE users SET username = $2 AND photourl = $3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: EditFirstNameOnly :exec
UPDATE users
SET
    firstname = $2
WHERE
    id = $1;

-- name: EditLastNameOnly :exec
UPDATE users
SET
    lastname = $2
WHERE
    id = $1;

-- name: EditBothNames :exec
UPDATE users
SET
    firstname = $2,
    lastname = $3
WHERE
    id = $1;