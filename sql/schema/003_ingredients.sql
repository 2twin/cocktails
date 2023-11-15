-- +goose Up
create table if not exists ingredients (
    id uuid primary key,
    name text unique not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
-- +goose Down
drop table ingredients;