-- USE 'glinteco_db'

CREATE TABLE roles
(
    id   int unique PRIMARY KEY NOT NULL,
    name varchar(64)
);

INSERT INTO roles VALUES (0, 'user');
INSERT INTO roles VALUES (1, 'admin');

CREATE TABLE users
(
    uuid         VARCHAR(36) UNIQUE PRIMARY KEY NOT NULL,
    username     varchar(128) UNIQUE            NOT NULL,
    display_name varchar(255)                   not null,
    email        varchar(255) unique            not null,
    role_id      int DEFAULT 0,
    active       boolean,
    created_at   datetime(3),
    updated_at   datetime(3),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);

CREATE TABLE contracts
(
    uuid      varchar(36) UNIQUE PRIMARY KEY NOT NULL,
    user_uuid varchar(36),
    name      varchar(255),
    info      blob,
    FOREIGN KEY (user_uuid) REFERENCES users (uuid)
)