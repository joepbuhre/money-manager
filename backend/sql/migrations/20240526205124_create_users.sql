-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  username varchar(256) not null,
  name varchar(256) not null,
  email varchar(256)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users
-- +goose StatementEnd
