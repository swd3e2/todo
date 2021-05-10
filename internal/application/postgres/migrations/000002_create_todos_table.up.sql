CREATE TABLE IF NOT EXISTS todos (
    id integer NOT NULL,
    user_id integer NOT NULL,
    text char(500),
    last_updated timestamp default current_timestamp,
    CONSTRAINT "pk_todos_id" PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
)