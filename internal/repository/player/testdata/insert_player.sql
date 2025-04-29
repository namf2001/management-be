TRUNCATE TABLE players RESTART IDENTITY CASCADE;
ALTER SEQUENCE players_id_seq RESTART WITH 1;

INSERT INTO players (id, department_id, full_name, position, jersey_number, date_of_birth, height_cm, weight_kg, phone, email, is_active)
VALUES (1, 1, 'Test Player', 'Forward', 10, '1990-01-01', 180, 75, '1234567890', 'test@example.com', true);

INSERT INTO players (id, department_id, full_name, position, jersey_number, date_of_birth, height_cm, weight_kg, phone, email, is_active)
VALUES (2, 2, 'Test Defender', 'Defender', 7, '1992-01-01', 175, 70, '0987654321', 'defender@example.com', true);