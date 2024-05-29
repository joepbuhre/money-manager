-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_token
ADD CONSTRAINT unique_user_token UNIQUE (token);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop CONSTRAINT unique_user_token;

-- +goose StatementEnd
