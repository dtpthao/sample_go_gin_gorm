CREATE TABLE users
(
    uuid       VARCHAR(36) UNIQUE PRIMARY KEY NOT NULL,
    username   varchar(128) UNIQUE            NOT NULL,
    password   varchar(255)                   not null,
    is_admin   bool DEFAULT false,
    active     boolean,
    created_at datetime(3),
    updated_at datetime(3)
);

CREATE TABLE contracts
(
    uuid      varchar(36) UNIQUE PRIMARY KEY NOT NULL,
    user_uuid varchar(36),
    name      varchar(255),
    info      blob,
    FOREIGN KEY (user_uuid) REFERENCES users (uuid)
)