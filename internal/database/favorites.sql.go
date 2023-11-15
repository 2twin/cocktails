// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: favorites.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addCocktailToFavorites = `-- name: AddCocktailToFavorites :one
insert into favorites_cocktails(favorite_id, cocktail_id)
values($1, $2)
returning favorite_id, cocktail_id
`

type AddCocktailToFavoritesParams struct {
	FavoriteID uuid.UUID
	CocktailID uuid.UUID
}

func (q *Queries) AddCocktailToFavorites(ctx context.Context, arg AddCocktailToFavoritesParams) (FavoritesCocktail, error) {
	row := q.db.QueryRowContext(ctx, addCocktailToFavorites, arg.FavoriteID, arg.CocktailID)
	var i FavoritesCocktail
	err := row.Scan(&i.FavoriteID, &i.CocktailID)
	return i, err
}

const createFavorite = `-- name: CreateFavorite :one
insert into favorites (id, user_id, created_at, updated_at)
values ($1, $2, $3, $4)
returning id, user_id, created_at, updated_at
`

type CreateFavoriteParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFavorite(ctx context.Context, arg CreateFavoriteParams) (Favorite, error) {
	row := q.db.QueryRowContext(ctx, createFavorite,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Favorite
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
