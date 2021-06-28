
CREATE TABLE update (
    update_id      TEXT      PRIMARY KEY NOT NULL,
    update_created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
