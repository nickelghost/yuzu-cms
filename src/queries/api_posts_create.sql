INSERT INTO
  posts (title, content, is_draft, created_at, updated_at)
VALUES
  ($1, $2, $3, NOW(), NOW())
RETURNING
  id, title, content, is_draft, created_at, updated_at
