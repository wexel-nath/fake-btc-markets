
CREATE TABLE IF NOT EXISTS market(
    market_id        TEXT PRIMARY KEY NOT NULL,
    base_asset       TEXT NOT NULL,
    quote_asset      TEXT NOT NULL,
    min_order_amount TEXT NOT NULL,
    max_order_amount TEXT NOT NULL,
    amount_decimals  TEXT NOT NULL,
    price_decimals   TEXT NOT NULL
);
