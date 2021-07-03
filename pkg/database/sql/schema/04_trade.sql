
CREATE TABLE IF NOT EXISTS trade(
    trade_id             SERIAL           PRIMARY KEY NOT NULL,
    order_id             INTEGER          NOT NULL REFERENCES "order"(order_id),
    trade_amount         NUMERIC(32, 16)  NOT NULL,
    trade_fee            NUMERIC(32, 16)  NOT NULL,
    trade_liquidity_type TEXT             NOT NULL,
    trade_created        TIMESTAMP        WITH TIME ZONE NOT NULL DEFAULT NOW()
);
