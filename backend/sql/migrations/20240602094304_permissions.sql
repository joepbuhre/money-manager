-- +goose Up
-- +goose StatementBegin
create table group_permissions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    group_id int not null,
    table_name varchar(256) not null,
    can_create boolean,
    can_read boolean,
    can_update boolean,
    can_delete boolean,
    object_id int
);
create table user_group (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int not null,
    group_id int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table group_permissions;
drop table user_group;
-- +goose StatementEnd
