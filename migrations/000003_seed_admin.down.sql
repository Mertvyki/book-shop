DELETE FROM bookshop.users WHERE email = 'admin@bookshop.ru';

ALTER TABLE bookshop.users DROP CONSTRAINT IF EXISTS chk_user_role;
ALTER TABLE bookshop.users ADD CONSTRAINT chk_user_role CHECK (role IN ('customer', 'admin'));

ALTER TABLE bookshop.users DROP CONSTRAINT IF EXISTS chk_full_name_length;
ALTER TABLE bookshop.users ADD CONSTRAINT chk_full_name_length CHECK (char_length(full_name) BETWEEN 15 AND 50);
