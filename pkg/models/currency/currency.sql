-- +goose Up
CREATE TABLE currency (
   name                               TEXT PRIMARY KEY,
   precision                          INT default 2 NOT NULL
);

-- +goose Down
DROP TABLE currency;