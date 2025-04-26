-- Xóa foreign key constraint với SET NULL
ALTER TABLE players DROP CONSTRAINT players_department_id_fkey;

-- Cập nhật tất cả records có department_id là NULL (nếu có)
UPDATE players SET department_id = 1 WHERE department_id IS NULL;

-- Thêm lại constraint NOT NULL cho department_id
ALTER TABLE players ALTER COLUMN department_id SET NOT NULL;

-- Thêm lại foreign key constraint như cũ
ALTER TABLE players
ADD CONSTRAINT players_department_id_fkey
FOREIGN KEY (department_id)
REFERENCES departments(id); 