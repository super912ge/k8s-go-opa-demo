CREATE TABLE IF NOT EXISTS "user"(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    email varchar unique not null
);

CREATE TABLE IF NOT EXISTS team(
    id serial PRIMARY KEY
);