ALTER DATABASE roles SET timezone TO 'UTC';

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO roles (name,created_at, updated_at) VALUES
('admin',now(),now()),
('staff',now(),now())
ON CONFLICT DO NOTHING;