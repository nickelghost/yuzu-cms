SELECT
  id, title, content, is_draft, created_at, updated_at
FROM
  posts
ORDER BY
  created_at
DESC
