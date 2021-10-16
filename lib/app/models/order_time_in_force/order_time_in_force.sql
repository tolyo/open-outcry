-- time constraints imposed on the order
CREATE TYPE order_time_in_force AS ENUM (
    -- good-til-cancelled (GTC) orders require a specific cancelling order, which can persist indefinitely
    -- this is the default state for all orders
    'GTC',

    -- an immediate or cancel (IOC) orders are immediately executed or cancelled by the exchange.
    -- Unlike FOK orders, IOC orders allow for partial fills.
    'IOC',

    -- Fill or kill (FOK) orders are usually limit orders that must be executed or cancelled immediately.
    -- Unlike IOC orders, FOK orders require the full quantity to be executed.
    'FOK',

    -- Good till date orders are valid until a certain date. The order gets
    -- canceled at the end of the day, if it is not executed.
    'GTD',

    -- Good till time orders are valid until a certain time with minute precision
    'GTT'
);