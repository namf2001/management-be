ALTER SEQUENCE match_players_id_seq RESTART WITH 1;
ALTER SEQUENCE players_id_seq RESTART WITH 1;
ALTER SEQUENCE departments_id_seq RESTART WITH 1;

-- Insert test department only if it doesn't exist
INSERT INTO departments (id, name) 
VALUES (1, 'Test Department')
ON CONFLICT (id) DO NOTHING;

INSERT INTO players (id, department_id, full_name, position, is_active) 
VALUES 
    (1, 1, 'Player One', 'Forward', true),
    (2, 1, 'Player Two', 'Midfielder', true),
    (3, 1, 'Player Three', 'Defender', true);

INSERT INTO match_players (id, match_id, player_id, minutes_played, goals_scored, assists, yellow_cards, red_card) 
VALUES 
    (1, 1, 1, 90, 2, 1, 0, false),
    (2, 1, 2, 85, 0, 2, 1, false),
    (3, 1, 3, 90, 0, 0, 0, true);