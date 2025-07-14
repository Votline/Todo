CREATE TABLE IF NOT EXISTS users (
    id            TEXT NOT NULL PRIMARY KEY,
    first_name    TEXT NOT NULL,
    last_name     TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    user_id     TEXT NOT NULL,
    task_id     SERIAL,
    title       TEXT NOT NULL,
    content     TEXT NOT NULL,
    category_id TEXT NOT NULL,
    done        BOOLEAN NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),

    PRIMARY KEY (user_id, task_id)
);
