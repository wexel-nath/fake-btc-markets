
CREATE TABLE IF NOT EXISTS market(
    market_id        TEXT            PRIMARY KEY NOT NULL,
    base_asset       TEXT            NOT NULL,
    quote_asset      TEXT            NOT NULL,
    min_order_amount NUMERIC(32, 16) NOT NULL,
    max_order_amount NUMERIC(32, 16) NOT NULL,
    amount_decimals  INTEGER         NOT NULL,
    price_decimals   INTEGER         NOT NULL
);
