-- +goose Up

CREATE TABLE trade_order(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    trading_account_id  BIGINT REFERENCES trading_account(id) NOT NULL,
    instrument_id       BIGINT REFERENCES instrument(id) NOT NULL,
    order_type          order_type NOT NULL,
    side                order_side NOT NULL,
    time_in_force       order_time_in_force NOT NULL,
    price               DECIMAL NOT NULL,
    amount              DECIMAL NOT NULL,
    open_amount         DECIMAL NOT NULL,
    status              trade_order_status NOT NULL default 'OPEN',
    updated_at          TIMESTAMP default current_timestamp NOT NULL,
    created_at          TIMESTAMP default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE trade_order CASCADE;
