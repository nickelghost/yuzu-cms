SELECT
  id, title, content, created_at, updated_at
FROM
  posts
WHERE
  id = $1
