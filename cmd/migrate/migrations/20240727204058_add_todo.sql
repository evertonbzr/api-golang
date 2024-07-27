-- +goose Up
CREATE TABLE todo (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  completed BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE todo;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
