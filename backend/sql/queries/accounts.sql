-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountByCode :one
SELECT * FROM accounts
WHERE code = $1;

-- name: Listaccounts :many
SELECT * FROM accounts
ORDER BY id;

-- name: CreateAccount :one
INSERT INTO accounts (code,name,description,account_type)
VALUES ($1, $2, $3, $4) 
RETURNING *;


