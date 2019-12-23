INSERT INTO
  posts (title, content, created_at, updated_at)
VALUES
  ($1, $2, NOW(), NOW())
RETURNING
  id, title, content, created_at, updated_at
