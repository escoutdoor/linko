-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores
(
    id uuid primary key default gen_random_uuid(),

    name text not null,
    description text,
    image_url text,
    address text not null,

    is_active boolean not null default true,

    location geography(Point, 4326) not null,

    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

CREATE INDEX idx_stores_location ON stores USING GIST (location);

CREATE TABLE IF NOT EXISTS categories
(
    id uuid primary key default gen_random_uuid(),
    store_id uuid not null references stores(id) on delete cascade,

    name text not null,

    position numeric not null default 0,

    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

CREATE TABLE IF NOT EXISTS products
(
    id uuid primary key default gen_random_uuid(),
    category_id uuid not null references categories(id) on delete cascade,

    name text not null,
    description text,
    image_url text,

    price_cents integer not null default 0,
    weight_grams int,

    is_available boolean default true,

    position numeric not null default 0,

    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS stores;
-- +goose StatementEnd
