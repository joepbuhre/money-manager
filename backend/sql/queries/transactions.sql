-- name: GetTransactions :one
SELECT * FROM transactions
WHERE id = $1;

-- name: GetTransactionsByAccount :one
SELECT * FROM transactions
WHERE account_id = $1;

-- name: Listtransactions :many
SELECT * FROM transactions
ORDER BY id;

