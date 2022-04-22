-- name: CreateAccount :one
INSERT INTO accounts (
  username, hashed_password, email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
SET hashed_password = $2,
    email = $3
WHERE id = $1
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;