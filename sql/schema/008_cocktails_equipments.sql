-- +goose Up
create table if not exists cocktails_equipments (
    cocktail_id uuid not null references cocktails(id) on delete cascade,
    equipment_id uuid not null references equipments(id) on delete cascade,
    constraint cocktails_equipments_pk primary key (cocktail_id, equipment_id)
);
-- +goose Down
drop table cocktails_equipments;