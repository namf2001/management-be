ALTER TABLE "matches"
ALTER COLUMN "our_score" SET NOT NULL,
ALTER COLUMN "our_score" SET DEFAULT 0;

ALTER TABLE "matches"
ALTER COLUMN "opponent_score" SET NOT NULL,
ALTER COLUMN "opponent_score" SET DEFAULT 0;