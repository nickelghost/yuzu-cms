CREATE TABLE pages(
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL,
  index INT NOT NULL,
  slug TEXT UNIQUE NOT NULL,
  in_navigation BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
