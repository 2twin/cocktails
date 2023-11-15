-- +goose Up
create table if not exists favorites (
    id uuid primary key,
    user_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null
);
-- +goose Down
drop table favorites;