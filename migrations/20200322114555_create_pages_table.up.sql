CREATE TABLE pages(
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL,
  index INT UNIQUE NOT NULL,
  slug TEXT UNIQUE NOT NULL,
  in_navigation BOOLEAN NOT NULL,
  heading TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
