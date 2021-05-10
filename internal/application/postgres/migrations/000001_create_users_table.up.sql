CREATE TABLE IF NOT EXISTS users (
    id integer NOT NULL,
    name char(50),
    last_name char(50),
    age integer,
    login char(20) NOT NULL,
    password char(500) NOT NULL,
    last_updated timestamp default current_timestamp,
    CONSTRAINT "pk_users_id" PRIMARY KEY (id)
);

create sequence users_id_seq
    start 1
    increment 1
    NO MAXVALUE
    CACHE 1;