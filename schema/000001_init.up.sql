CREATE TABLE users
(
    id            SERIAL UNIQUE NOT NULL ,
    name          VARCHAR(255)  NOT NULL,
    username      VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL
);

CREATE TABLE todo_lists
(
    id             SERIAL NOT NULL UNIQUE ,
    title          VARCHAR(255) NOT NULL,
    description    VARCHAR(255)
);

CREATE TABLE todo_items
(
    id              SERIAL PRIMARY KEY UNIQUE ,
    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(255),
    done            BOOLEAN      NOT NULL DEFAULT false
);

CREATE TABLE users_lists
(
    id          SERIAL NOT NULL UNIQUE,
    user_id     INT  REFERENCES users(id) ON DELETE CASCADE NOT NULL ,
    list_id     INT  REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE lists_items
(
    id          SERIAL NOT NULL UNIQUE,
    list_id     INTEGER NOT NULL REFERENCES todo_lists(id) ON DELETE CASCADE,
    item_id     INTEGER NOT NULL REFERENCES todo_items(id) On DELETE CASCADE
);

