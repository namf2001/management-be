ALTER SEQUENCE departments_id_seq RESTART WITH 1;

INSERT INTO departments ( name, description)
VALUES
    ( 'Test Department', 'This is a test department');