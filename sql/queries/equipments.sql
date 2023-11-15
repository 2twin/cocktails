-- name: CreateEquipment :one
insert into equipments (id, name, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;
-- name: GetEquipments :many
select *
from equipments;
-- name: GetEquipmentById :one
select *
from equipments
where id = $1;
-- name: GetCocktailEquipments :many
select *
from cocktails_equipments
where cocktail_id = $1;