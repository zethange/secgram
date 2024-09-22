-- migrate:up
CREATE TABLE messages(
    id serial primary key,
    content text not null,
    created_at timestamp not null default now(),
    user_id serial not null,
    chat_id serial not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chats(id)
);

-- migrate:down
DROP TABLE messages;

