UPDATE
  posts
SET
  title = $2,
  content = $3,
  updated_at = NOW()
WHERE
  id = $1
RETURNING
  id, title, content, created_at, updated_at
