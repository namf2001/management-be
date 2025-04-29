TRUNCATE TABLE departments RESTART IDENTITY CASCADE;
ALTER SEQUENCE departments_id_seq RESTART WITH 1;

INSERT INTO departments (id, name, description)
VALUES 
    (1, 'Test Department 1', 'This is test department 1'),
    (2, 'Test Department 2', 'This is test department 2');