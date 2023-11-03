SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."board_action"
(
    id           SERIAL PRIMARY KEY,
    cafe_id      int     not null,
    cafe_type_id int     not null,
    read_roles   varchar not null,
    create_roles varchar not null,
    update_roles varchar not null,
    update_able  bool default false,
    delete_roles varchar not null,
    created_at   timestamptz
);


create unique index ba_cafe_id_cafe_type_id_unique on board_action (cafe_id, cafe_type_id);
