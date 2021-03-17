CREATE TABLE IF NOT EXISTS colors(
id SERIAL PRIMARY KEY,
name text NOT NULL,
hexadecimal text,
R int,
G int,
B int,
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);