SET statement_timeout = 0;

CREATE TABLE user_entities (
    id VARCHAR DEFAULT gen_random_uuid() PRIMARY KEY,
    firstname VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM NOW()),
    deleted_at TIMESTAMP NULL
);
