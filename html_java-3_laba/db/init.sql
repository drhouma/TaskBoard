CREATE TABLE IF NOT EXISTS users
(
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS comments
(
    id        SERIAL PRIMARY KEY,
    user_name TEXT   NOT NULL,
    message   TEXT   NOT NULL,

    FOREIGN KEY (user_name) REFERENCES users (name)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
