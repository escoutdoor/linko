-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id uuid primary key default gen_random_uuid(),
    first_name text,
    last_name text,
    email text unique,
    phone_number text unique,
    password text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
