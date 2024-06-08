-- +goose Up
-- +goose StatementBegin
create table transactions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    account_id int,
    description varchar(500),
    debit decimal(19,4),
    credit decimal(19,4),
    reconciled boolean default FALSE
);

create table accounts (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    code varchar(50) not null,
    name varchar(256) not null,
    description varchar(500),
    account_type int not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table transactions;
drop table accounts;
-- +goose StatementEnd
