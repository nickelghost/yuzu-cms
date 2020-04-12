UPDATE
  posts
SET
  title = $2,
  content = $3,
  is_draft = $4,
  updated_at = NOW()
WHERE
  id = $1
RETURNING
  id, title, content, is_draft, created_at, updated_at
