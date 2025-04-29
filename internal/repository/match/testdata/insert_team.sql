ALTER SEQUENCE teams_id_seq RESTART WITH 1;
-- Insert test data for teams
INSERT INTO teams (id, name, company_name, contact_person, contact_email) 
VALUES 
    (1, 'Test Team 1', 'Test Company 1', 'Test Contact 1', 'test1@example.com'),
    (2, 'Test Team 2', 'Test Company 2', 'Test Contact 2', 'test2@example.com');