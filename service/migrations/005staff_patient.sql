CREATE TABLE IF NOT EXISTS staff_patients (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    patient_id INTEGER NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_staff_patient_user
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_staff_patient_patient
        FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE,

    CONSTRAINT uq_staff_patient UNIQUE (user_id, patient_id)
);

INSERT INTO staff_patients (id, user_id, patient_id,created_at, updated_at)
VALUES
(1, 2, 1,now(),now()), -- staff1 ดูแล patient 1
(2, 2, 2,now(),now()) -- staff1 ดูแล patient 2
ON CONFLICT DO NOTHING;