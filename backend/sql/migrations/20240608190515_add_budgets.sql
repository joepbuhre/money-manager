-- +goose Up
-- +goose StatementBegin
create table budgets (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    name int,
    description varchar(500),
    budget_amount decimal(19,4) not null default 0,
    budget_type int not null default 1
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table budgets
-- +goose StatementEnd
