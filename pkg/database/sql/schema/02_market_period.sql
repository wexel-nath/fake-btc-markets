
CREATE TABLE IF NOT EXISTS market_period(
    market_period_id  SERIAL          PRIMARY KEY NOT NULL,
    market_id         TEXT            NOT NULL REFERENCES market(market_id),
    time_period_start TIMESTAMP       WITH TIME ZONE NOT NULL,
    time_period_end   TIMESTAMP       WITH TIME ZONE NOT NULL,
    price_open        NUMERIC(32, 16) NOT NULL,
    price_high        NUMERIC(32, 16) NOT NULL,
    price_low         NUMERIC(32, 16) NOT NULL,
    price_close       NUMERIC(32, 16) NOT NULL,
    volume_traded     NUMERIC(32, 16) NOT NULL,

    CONSTRAINT market_period_unique UNIQUE (market_id, time_period_start, time_period_end)
);
