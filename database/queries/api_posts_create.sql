INSERT INTO
  posts (title, content, created_at, updated_at)
VALUES
  ($1, $2, $3, $4)
RETURNING
  id, title, content, created_at, updated_at
