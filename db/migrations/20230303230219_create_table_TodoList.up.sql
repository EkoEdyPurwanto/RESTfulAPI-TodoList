CREATE TABLE TodoList(
                         id          int                      NOT NULL PRIMARY KEY AUTO_INCREMENT,
                         title       varchar(50)              NOT NULL,
                         description text                     NOT NULL,
                         status      enum ('PENDING', 'DONE') NOT NULL DEFAULT 'PENDING',
                         created_at  timestamp                NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at  timestamp                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
