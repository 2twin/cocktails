-- name: CreateUser :one
insert into users (id, created_at, updated_at, name, api_key)
values (
        $1,
        $2,
        $3,
        $4,
        encode(sha256(random()::text::bytea), 'hex')
    )
returning *;
-- name: GetUserByAPIKey :one
select *
from users
where api_key = $1;
-- name: GetFavoritesByUser :one
select *
from favorites
where user_id = $1;