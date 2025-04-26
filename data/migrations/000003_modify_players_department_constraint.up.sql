-- Cho phép department_id được NULL
ALTER TABLE players ALTER COLUMN department_id DROP NOT NULL;

-- Xóa foreign key constraint cũ
ALTER TABLE players DROP CONSTRAINT players_department_id_fkey;

-- Thêm foreign key constraint mới với ON DELETE SET NULL
ALTER TABLE players 
ADD CONSTRAINT players_department_id_fkey 
FOREIGN KEY (department_id) 
REFERENCES departments(id) 
ON DELETE SET NULL; 