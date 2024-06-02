-- +goose Up
-- +goose StatementBegin
alter table users add is_superadmin boolean default FALSE
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users drop is_superadmin
-- +goose StatementEnd
