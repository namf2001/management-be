ALTER TABLE players DROP CONSTRAINT players_department_id_fkey;

UPDATE players
SET department_id = (SELECT id FROM departments LIMIT 1)
WHERE department_id IS NULL;

ALTER TABLE players ALTER COLUMN department_id SET NOT NULL;

ALTER TABLE players
ADD CONSTRAINT players_department_id_fkey
FOREIGN KEY (department_id)
REFERENCES departments(id); 