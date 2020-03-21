SELECT
  id, title, content, is_draft, created_at, updated_at
FROM
  posts
WHERE
  NOT $1 OR is_draft = TRUE
ORDER BY
  created_at
DESC
