-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (username, name, email)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetUserFromToken :one
select u.* 
from users u
inner join user_token ut on u.id = ut.user_id and ut.token = $1;

-- name: AddToken :exec
insert into user_token (user_id, token)
values ($1, $2)
ON CONFLICT (token) DO NOTHING;

-- name: LoginUser :one
insert into user_token (user_id, token)
values($1, $2)
ON CONFLICT (token) DO NOTHING
RETURNING *;