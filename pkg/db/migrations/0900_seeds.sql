-- +goose Up

INSERT INTO currency(name, precision)
VALUES ('EUR',2),
       ('BTC',5);

INSERT INTO app_entity (pub_id, external_id, type)
VALUES ('MASTER','MASTER','MASTER');

SELECT create_payment_account('MASTER','EUR');
SELECT create_payment_account('MASTER','BTC');

INSERT INTO instrument(name, base_currency, quote_currency, fx_instrument)
VALUES ('BTC_EUR', 'BTC', 'EUR', TRUE);

INSERT INTO instrument(name, quote_currency)
VALUES ('SPX', 'EUR');

-- INSERT INTO fee(type, currency_name, min)
-- VALUES ('DEPOSIT_FEE', 'EUR', 1.00);

-- +goose Down
