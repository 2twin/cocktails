-- name: CreateCocktail :one
insert into cocktails (
        id,
        name,
        instructions,
        created_at,
        updated_at
    )
values ($1, $2, $3, $4, $5)
returning *;
-- name: GetCocktails :many
select *
from cocktails;
-- name: GetCocktailById :one
select *
from cocktails
where id = $1;
-- name: AddEquipmentToCocktail :one
insert into cocktails_equipments (cocktail_id, equipment_id)
values ($1, $2)
returning *;
-- name: AddIngredientToCocktail :one
insert into cocktails_ingredients (cocktail_id, ingredient_id)
values($1, $2)
returning *;
-- name: GetCocktailsByEquipment :many
select cocktail_id
from cocktails_equipments
where equipment_id = $1;
-- name: GetCocktailsByIngredient :many
select cocktail_id
from cocktails_ingredients
where ingredient_id = $1;