-- name: CreateFavorite :one
insert into favorites (id, user_id, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;
-- name: AddCocktailToFavorites :one
insert into favorites_cocktails(favorite_id, cocktail_id)
values($1, $2)
returning *;