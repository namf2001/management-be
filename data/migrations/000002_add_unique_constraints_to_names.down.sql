-- Remove unique constraints from name fields


-- Remove unique constraint from player full names
ALTER TABLE "players" DROP CONSTRAINT IF EXISTS "unique_player_full_name";

-- Remove unique constraint from team names
ALTER TABLE "teams" DROP CONSTRAINT IF EXISTS "unique_team_name";

-- Remove unique constraint from department names
ALTER TABLE "departments" DROP CONSTRAINT IF EXISTS "unique_department_name";

-- Remove unique constraint from usernames
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "unique_username";