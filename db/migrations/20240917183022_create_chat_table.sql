-- migrate:up
CREATE TABLE chats (
    id serial primary key,
    type varchar(255) not null default 'private',
    created_at timestamp not null default now()
);
-- migrate:down
DROP TABLE chats;

