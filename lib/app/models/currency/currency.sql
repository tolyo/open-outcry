CREATE TABLE IF NOT EXISTS currency (
       name                               TEXT PRIMARY KEY,
       precision                          INT default 2 NOT NULL,
       updated_at                         TIMESTAMP DEFAULT current_timestamp NOT NULL,
       created_at                         TIMESTAMP DEFAULT current_timestamp NOT NULL
);