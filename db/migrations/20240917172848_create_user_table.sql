-- migrate:up
CREATE TABLE users(
    id serial primary key,
    full_name varchar(255) not null,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp default now() not null,
    last_seen timestamp default now() not null
);

-- migrate:down
DROP TABLE users;
