
CREATE TABLE IF NOT EXISTS "order"(
    order_id             SERIAL           PRIMARY KEY NOT NULL,
    market_id            TEXT             NOT NULL REFERENCES market(market_id),
    order_price          NUMERIC(32, 16)  NOT NULL,
    order_amount         NUMERIC(32, 16)  NOT NULL,
    order_type           TEXT             NOT NULL,
    order_side           TEXT             NOT NULL,
    order_trigger_price  NUMERIC(32, 16),
    order_trigger_amount NUMERIC(32, 16),
    order_time_in_force  TEXT             NOT NULL DEFAULT 'GTC',
    order_post_only      BOOLEAN          NOT NULL DEFAULT FALSE,
    order_self_trade     TEXT             NOT NULL DEFAULT 'A',
    order_created        TIMESTAMP        WITH TIME ZONE NOT NULL DEFAULT NOW(),
    order_status         TEXT             NOT NULL DEFAULT 'Accepted',
    client_order_id      TEXT
);
