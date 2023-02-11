-- migrate:up
CREATE TABLE IF NOT EXISTS users (
  id UUID primary key DEFAULT uuid_generate_v4() NOT NULL,
  email text,
  username text,
  first_name text,
  last_name text,
  password text,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  deleted_at timestamp
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users(email);

-- migrate:down
DROP INDEX IF EXISTS users_email_idx;
DROP TABLE IF EXISTS users;
