ALTER DATABASE genders SET timezone TO 'UTC';

CREATE TABLE IF NOT EXISTS genders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) NOT NULL UNIQUE, -- // Male , Female
    abbreviation VARCHAR(10) NOT NULL UNIQUE, -- M, F
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO genders (name,abbreviation,created_at, updated_at) VALUES
('Male','M',now(),now()),
('Female','F',now(),now())
ON CONFLICT DO NOTHING;
