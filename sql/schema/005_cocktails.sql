-- +goose Up
create table if not exists cocktails (
    id uuid primary key,
    name text not null,
    instructions text [] not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
-- +goose Down
drop table cocktails;