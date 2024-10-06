-- +goose Up
-- Model of states that a trade order can undergo during its
-- lifecycle. Transitional states hold funds in `payment_account.amount_reserved`
-- Final states should release all reserved funds.
CREATE TYPE trade_order_status AS ENUM(
    -- TRANSITIONAL STATES
    -- order remains in the order book and has no fills
    'OPEN',
    -- order remains in the order book but with partial fill
    'PARTIALLY_FILLED',
    -- FINAL STATES
    -- order has been cancelled by the user without any fills
    'CANCELLED',
    -- order has been partially filled but then cancelled by user
    'PARTIALLY_CANCELLED',
    -- order has been partially filled but then rejected.
    -- Possible rejection reasons:
    --  * IOC
    --  * GTD
    --  * GTT
    'PARTIALLY_REJECTED',
    -- order has been filled and removed from the order book
    'FILLED',
    -- order rejected from the order book.
    -- Possible rejection reasons:
    --  * insufficiency of funds
    --  * FOK or IOC order that cannot be filled
    'REJECTED'
);

-- +goose Down
DROP TYPE trade_order_status CASCADE;

