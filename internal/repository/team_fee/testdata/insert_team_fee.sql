TRUNCATE TABLE team_fees CASCADE;
ALTER SEQUENCE team_fees_id_seq RESTART WITH 1;

-- Insert test data
INSERT INTO team_fees (id, amount, payment_date, description, created_at, updated_at) VALUES
(1, 100.50, '2023-10-01', 'Monthly team fee', NOW(), NOW()),
(2, 200.75, '2023-10-02', 'Advance payment', NOW(), NOW());
