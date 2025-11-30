-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS drivers
(
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null unique,

    status int not null default 1,

    review_count integer default 0,
    total_rating_sum decimal default 0.0,

    vehicle_type int,
    vehicle_model text,
    vehicle_plate_number text,
    vehicle_color text,

    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS drivers;

-- +goose StatementEnd
