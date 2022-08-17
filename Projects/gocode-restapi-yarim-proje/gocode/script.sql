CREATE TABLE books (
    id          VARCHAR PRIMARY KEY,
    title       VARCHAR NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
)