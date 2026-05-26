-- +goose Up
CREATE TYPE component_type AS ENUM ('tenant', 'third-party');

CREATE TABLE components (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    team TEXT NOT NULL,
    area TEXT NOT NULL,
    type component_type NOT NULL,
    UNIQUE (name, team, area)
);

-- +goose Down
DROP TABLE components;
DROP TYPE component_type;