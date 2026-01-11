ALTER DATABASE patients SET timezone TO 'UTC';


CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,

    first_name_th VARCHAR(100),
    middle_name_th VARCHAR(100),
    last_name_th VARCHAR(100),

    first_name_en VARCHAR(100),
    middle_name_en VARCHAR(100),
    last_name_en VARCHAR(100),

    date_of_birth TIMESTAMPTZ DEFAULT NOW(),
    patient_hn VARCHAR(50) NOT NULL UNIQUE,
    national_id VARCHAR(20),
    passport_id VARCHAR(20),

    phone_number VARCHAR(20),
    email VARCHAR(100),

    gender_id INTEGER NOT NULL,
    hospital_id INTEGER NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO patients (
    id, first_name_th, middle_name_th, last_name_th,
    first_name_en, middle_name_en, last_name_en,
    date_of_birth, patient_hn,
    national_id, passport_id, phone_number,
    email, gender_id, hospital_id,
    created_at, updated_at
)
VALUES
(1, 'สมชาย','-', 'ใจดี', 'Somchai','-', 'Jaidee',
    '1990-05-12 00:00:00+07', 'HN0001', '1103701234567','123456789', '0812345678', 'somchai@example.com',
    1,1, now(), now()),
(2, 'สมหญิง','-', 'รักสุข', 'Somying','-', 'Raksuk',
    '1990-09-20 00:00:00+07', 'HN0002', '1103707654321','567875890', '0898765432', 'somying@example.com',
    2,1, now(), now())
ON CONFLICT DO NOTHING;