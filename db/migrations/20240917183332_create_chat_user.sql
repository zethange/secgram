-- migrate:up
CREATE TABLE chats_users (
    id serial primary key,
    chat_id serial not null,
    user_id serial not null,
    FOREIGN KEY (chat_id) REFERENCES chats(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);


-- migrate:down
DROP TABLE chats_users;
