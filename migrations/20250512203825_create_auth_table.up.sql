CREATE TABLE auth
(
    id      SERIAL PRIMARY KEY,
    user_id INT,
    session VARCHAR(225),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

