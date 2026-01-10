ALTER DATABASE hospitals SET timezone TO 'UTC';


CREATE TABLE IF NOT EXISTS hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


INSERT INTO hospitals (name,created_at, updated_at) VALUES
('Bangkok Hospital',now(),now())
ON CONFLICT DO NOTHING;