CREATE TABLE IF NOT EXISTS currency (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
       name                               TEXT UNIQUE NOT NULL,
       precision                          INT default 2 NOT NULL,
       updated_at                         TIMESTAMP DEFAULT current_timestamp NOT NULL,
       created_at                         TIMESTAMP DEFAULT current_timestamp NOT NULL
);