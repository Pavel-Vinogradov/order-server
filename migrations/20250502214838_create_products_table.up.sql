CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(225),
    description TEXT,
    images      TEXT[],
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
