-- +goose Up
create table if not exists favorites_cocktails (
    favorite_id uuid not null references favorites(id) on delete cascade,
    cocktail_id uuid not null references cocktails(id) on delete cascade,
    constraint favorites_cocktails_pk primary key (favorite_id, cocktail_id)
);
-- +goose Down
drop table favorites_cocktails;