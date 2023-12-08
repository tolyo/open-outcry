-- +goose Up

INSERT INTO currency(name, precision)
VALUES ('EUR',2),
       ('USD',2),
       ('BTC',5);

INSERT INTO app_entity (pub_id, external_id, type)
VALUES ('MASTER','MASTER','MASTER');

SELECT create_payment_account('MASTER','EUR');
SELECT create_payment_account('MASTER','BTC');

INSERT INTO instrument(name, base_currency, quote_currency, fx_instrument)
VALUES ('BTC_EUR', 'BTC', 'EUR', TRUE);

INSERT INTO instrument(name, quote_currency)
VALUES ('SPX', 'EUR');

-- +goose Down
