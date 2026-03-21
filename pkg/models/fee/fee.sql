-- +goose Up
CREATE TABLE IF NOT EXISTS fee (
     type                               TEXT NOT NULL,
     currency_name                      TEXT REFERENCES currency(name) NOT NULL,
     min                                NUMERIC NULL,
     max                                NUMERIC NULL,
     percentage                         NUMERIC NULL,
     updated_at                         TIMESTAMPTZ default current_timestamp NOT NULL,
     created_at                         TIMESTAMPTZ default current_timestamp NOT NULL,
     PRIMARY KEY(type, currency_name)
);

-- +goose Down
DROP TABLE fee CASCADE;
