CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(225),
    password   VARCHAR(225),
    email      varchar(20),
    phone      VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
