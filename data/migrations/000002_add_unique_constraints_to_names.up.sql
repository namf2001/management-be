-- Add unique constraints to name fields

-- Make username unique
ALTER TABLE "users" ADD CONSTRAINT "unique_username" UNIQUE ("username");

-- Make department names unique
ALTER TABLE "departments" ADD CONSTRAINT "unique_department_name" UNIQUE ("name");

-- Make team names unique
ALTER TABLE "teams" ADD CONSTRAINT "unique_team_name" UNIQUE ("name");

-- Make player full names unique
ALTER TABLE "players" ADD CONSTRAINT "unique_player_full_name" UNIQUE ("full_name");