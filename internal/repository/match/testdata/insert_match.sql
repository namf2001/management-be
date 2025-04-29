

-- Insert test data for matches
INSERT INTO matches (id, opponent_team_id, match_date, venue, is_home_game, status, notes)
VALUES
    (1, 1, '2025-04-30 15:00:00', 'Test Venue 1', true, 'scheduled', 'Test Notes 1'),
    (2, 2, '2025-05-01 18:00:00', 'Test Venue 2', false, 'scheduled', 'Test Notes 2');