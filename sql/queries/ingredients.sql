-- name: CreateIngredient :one
insert into ingredients (id, name, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;
-- name: GetIngredients :many
select *
from ingredients;
-- name: GetIngredientById :one
select *
from ingredients
where id = $1;
-- name: GetCocktailIngredients :many
select *
from cocktails_ingredients
where cocktail_id = $1;