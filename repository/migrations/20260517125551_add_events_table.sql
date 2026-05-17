-- +goose Up
CREATE TABLE events (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    concurrency_target REAL NOT NULL
);

-- +goose Down
DROP TABLE events;
