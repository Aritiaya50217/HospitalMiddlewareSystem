-- สร้างตาราง role
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- สร้างตาราง hospital
CREATE TABLE IF NOT EXISTS hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- สร้างตาราง user
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    hospital_id INT NOT NULL REFERENCES hospitals(id),
    role_id INT NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- สร้างตัวอย่าง role
INSERT INTO roles (name,created_at, updated_at) VALUES
('admin',now(),now()),
('staff',now(),now())
ON CONFLICT DO NOTHING;

-- สร้างตัวอย่าง hospital
INSERT INTO hospitals (name,created_at, updated_at) VALUES
('Bangkok Hospital',now(),now()),
('Samitivej Hospital',now(),now())
ON CONFLICT DO NOTHING;

-- สร้างตัวอย่าง user (password hash: admin1234 , staff1234)
INSERT INTO users (username, password, hospital_id, role_id,created_at, updated_at)
VALUES 
('admin', '$2a$10$VQG3Sk0DsVW/qn7HXMF6nevPl0baoxx1XnoIvlh2qQSBNXlIUudeu', 1, 1 ,now(),now()),
('staff1', '$2a$10$Kd.f60dC6nJUnGebiYMDW.g8b/8JTMavYN9S6Ni7Kpjfxz70aD3le', 1, 2, now(),now())
ON CONFLICT DO NOTHING;
