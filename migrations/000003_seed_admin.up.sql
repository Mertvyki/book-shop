ALTER TABLE bookshop.users DROP CONSTRAINT IF EXISTS chk_full_name_length;
ALTER TABLE bookshop.users ADD CONSTRAINT chk_full_name_length CHECK (char_length(full_name) BETWEEN 2 AND 50);

ALTER TABLE bookshop.users DROP CONSTRAINT IF EXISTS chk_user_role;
ALTER TABLE bookshop.users ADD CONSTRAINT chk_user_role CHECK (role IN ('customer', 'admin', 'employee'));

INSERT INTO bookshop.users (email, password_hash, full_name, role)
VALUES (
    'admin@bookshop.ru',
    '$2a$12$XQ1xVie/Zn0FqcXuJ/wvIORUg.U48YYi7XgWJqKQhIIAw7k.ajNie',
    'Administrator',
    'admin'
);
