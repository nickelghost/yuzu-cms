SELECT
  id, title, content, created_at, updated_at
FROM
  posts
WHERE
  is_draft = FALSE
ORDER BY
  created_at
DESC
