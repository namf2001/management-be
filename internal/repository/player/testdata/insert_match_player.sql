-- Ensure we have a team and match first
INSERT INTO teams (id, name, company_name) 
VALUES (1, 'Test Team', 'Test Company') 
ON CONFLICT (id) DO NOTHING;

INSERT INTO matches (id, opponent_team_id, match_date, venue, is_home_game)
VALUES (1, 1, '2025-01-01', 'Test Venue', true)
ON CONFLICT (id) DO NOTHING;

-- Insert match_player record
INSERT INTO match_players (match_id, player_id, minutes_played, goals_scored, assists, yellow_cards, red_card)
VALUES (1, 1, 90, 1, 0, 0, false);