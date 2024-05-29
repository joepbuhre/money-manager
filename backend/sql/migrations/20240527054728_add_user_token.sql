-- +goose Up
-- +goose StatementBegin
create table user_token (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    token varchar(256) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_token
-- +goose StatementEnd
