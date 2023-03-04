CREATE TYPE todo_status AS ENUM ('PENDING', 'DONE');

CREATE TABLE TodoList (
                          id          SERIAL PRIMARY KEY,
                          title       VARCHAR(50) NOT NULL,
                          description TEXT NOT NULL,
                          status      todo_status NOT NULL DEFAULT 'PENDING',
                          created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);