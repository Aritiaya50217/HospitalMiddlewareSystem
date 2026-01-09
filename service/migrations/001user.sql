CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    hospital_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_users_hospital
        FOREIGN KEY (hospital_id) REFERENCES hospitals(id),

    CONSTRAINT fk_users_role
        FOREIGN KEY (role_id) REFERENCES roles(id)
);

INSERT INTO users (username, password, hospital_id, role_id,created_at, updated_at)
VALUES 
('admin', '$2a$10$VQG3Sk0DsVW/qn7HXMF6nevPl0baoxx1XnoIvlh2qQSBNXlIUudeu', 1, 1,now(),now()), -- password : admin1234 
('staff1', '$2a$10$Kd.f60dC6nJUnGebiYMDW.g8b/8JTMavYN9S6Ni7Kpjfxz70aD3le', 1, 2,now(),now())  -- password : staff1234
ON CONFLICT DO NOTHING;