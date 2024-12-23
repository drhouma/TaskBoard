CREATE TABLE IF NOT EXISTS users
(
    "name"     TEXT  PRIMARY KEY,
    "password" TEXT  NOT NULL
);

CREATE TABLE IF NOT EXISTS categories
(
    "name" TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS tasks
(
    "id"          SERIAL PRIMARY KEY,
    "user"        TEXT   NOT NULL,
    "description" TEXT   NOT NULL,
    "category"    TEXT   NOT NULL,

    FOREIGN KEY ("user")     REFERENCES users ("name"),
    FOREIGN KEY ("category") REFERENCES categories ("name"),

    UNIQUE ("user", "description")
);

INSERT INTO users VALUES ('Andrey', '123456');
INSERT INTO users VALUES ('Aboba', '654321');
INSERT INTO users VALUES ('Kristina', 'qwerty');

INSERT INTO categories VALUES ('todo');
INSERT INTO categories VALUES ('in progress');
INSERT INTO categories VALUES ('done');

INSERT INTO tasks VALUES (DEFAULT, 'Andrey', 'Лабы по ммус', 'todo');
INSERT INTO tasks VALUES (DEFAULT, 'Andrey', 'Проект по червям', 'in progress');
INSERT INTO tasks VALUES (DEFAULT, 'Andrey', 'API golang', 'done');

INSERT INTO tasks VALUES (DEFAULT, 'Aboba', 'Смонтировать видео', 'todo');
INSERT INTO tasks VALUES (DEFAULT, 'Aboba', 'Снять видео', 'in progress');
INSERT INTO tasks VALUES (DEFAULT, 'Aboba', 'Написать сценарий к видео', 'done');

INSERT INTO tasks VALUES (DEFAULT, 'Kristina', 'Найти работу', 'todo');
INSERT INTO tasks VALUES (DEFAULT, 'Kristina', 'Приготовить ужин', 'in progress');
INSERT INTO tasks VALUES (DEFAULT, 'Kristina', 'Сходить по магазинам', 'done');
