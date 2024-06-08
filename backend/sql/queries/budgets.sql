-- name: GetBudgetById :one
SELECT * FROM budgets
WHERE id = $1;

-- name: GetBudgetByName :one
SELECT * FROM budgets
WHERE name = $1;

-- name: ListBudgets :many
SELECT * FROM budgets
ORDER BY id;

-- name: CreateBudget :one
INSERT INTO budgets (name, description, budget_amount, budget_type)
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: DeleteBudget :exec
delete from budgets where id = $1;