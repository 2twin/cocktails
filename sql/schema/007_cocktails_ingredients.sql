-- +goose Up
create table if not exists cocktails_ingredients (
    cocktail_id uuid not null references cocktails(id) on delete cascade,
    ingredient_id uuid not null references ingredients(id) on delete cascade,
    constraint cocktails_ingredients_pk primary key (cocktail_id, ingredient_id)
);
-- +goose Down
drop table cocktails_ingredients;