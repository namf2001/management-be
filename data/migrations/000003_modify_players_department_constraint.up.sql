ALTER TABLE players ALTER COLUMN department_id DROP NOT NULL;

ALTER TABLE players DROP CONSTRAINT players_department_id_fkey;

ALTER TABLE players
ADD CONSTRAINT players_department_id_fkey 
FOREIGN KEY (department_id) 
REFERENCES departments(id) 
ON DELETE SET NULL; 