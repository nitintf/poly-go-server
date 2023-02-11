-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

-- migrate:down
DROP EXTENSION IF EXISTS "uuid-ossp";
