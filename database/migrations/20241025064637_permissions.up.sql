-- SET statement_timeout = 0;

-- CREATE TABLE permissions (
--     id VARCHAR DEFAULT gen_random_uuid() PRIMARY KEY,
--     group_name VARCHAR NOT NULL,
--     name VARCHAR NOT NULL,
--     is_active BOOL NOT NULL,
--     created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
--     updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
--     deleted_at TIMESTAMP NULL
-- );