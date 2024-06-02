-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (username, name, email)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: DeleteUser :exec
delete from users where id = $1;

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

-- name: GetUserPermission :one
select u.is_superadmin, gp.can_create, gp.can_read, gp.can_update, gp.can_delete, gp.object_id
from users u
left join user_group ug on u.id = ug.user_id
left join group_permissions gp on ug.group_id = gp.group_id
where 
    (table_name = $1 or u.is_superadmin = TRUE)
    and u.id = $2;