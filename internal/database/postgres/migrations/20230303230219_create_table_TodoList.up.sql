CREATE TABLE users
(
    user_id    SERIAL PRIMARY KEY,
    username   VARCHAR(50) UNIQUE  NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE todo_status AS ENUM ('PENDING', 'DONE');

CREATE TABLE TodoList
(
    todo_id     SERIAL PRIMARY KEY,
    user_id     INTEGER     NOT NULL,
    title       VARCHAR(50) NOT NULL,
    description TEXT        NOT NULL,
    status      todo_status NOT NULL DEFAULT 'PENDING',
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);