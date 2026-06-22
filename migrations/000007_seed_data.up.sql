-- Seed data for demonstration

-- Publishers
INSERT INTO bookshop.publishers (name) VALUES
    ('Эксмо'),
    ('АСТ'),
    ('МИФ'),
    ('Питер');

-- Categories
INSERT INTO bookshop.categories (name, slug, description) VALUES
    ('Художественная литература', 'fiction', 'Романы, повести, рассказы'),
    ('Научная литература', 'non-fiction', 'Научно-популярные книги'),
    ('Фантастика', 'sci-fi', 'Научная фантастика и фэнтези'),
    ('История', 'history', 'Исторические книги'),
    ('Технологии', 'tech', 'Книги о программировании и IT'),
    ('Детская литература', 'children', 'Книги для детей');

-- Authors
INSERT INTO bookshop.authors (name, bio, birth_year) VALUES
    ('Фёдор Достоевский', 'Великий русский писатель и мыслитель', 1821),
    ('Лев Толстой', 'Один из наиболее известных русских писателей', 1828),
    ('Михаил Булгаков', 'Русский писатель и драматург', 1891),
    ('Братья Стругацкие', 'Советские писатели-фантасты', NULL),
    ('Стивен Кинг', 'Американский писатель в жанре ужасов', 1947),
    ('Роберт Мартин', 'Автор книг по программированию', 1952),
    ('Юваль Ной Харари', 'Израильский историк и писатель', 1976),
    ('Джоан Роулинг', 'Британская писательница', 1965);

-- Books
INSERT INTO bookshop.books (title, description, isbn, price, book_type, stock_quantity, file_key, cover_image_key, publisher_id) VALUES
    ('Преступление и наказание', 'Роман о нравственных исканиях и социальных противоречиях', '978-5-04-123456-1', 599.00, 'physical', 10, NULL, NULL, 1),
    ('Война и мир. Том 1', 'Эпопея о жизни русского общества в эпоху наполеоновских войн', '978-5-04-123456-2', 799.00, 'physical', 5, NULL, NULL, 1),
    ('Мастер и Маргарита', 'Мистический роман о любви и силе творчества', '978-5-04-123456-3', 450.00, 'digital', NULL, 'books/master-and-margarita.epub', NULL, 1),
    ('Пикник на обочине', 'Знаменитый роман братьев Стругацких о сталкерах', '978-5-04-123456-4', 350.00, 'physical', 8, NULL, NULL, 2),
    ('Оно', 'Эпический роман ужасов о противостоянии детей и древнего зла', '978-5-04-123456-5', 899.00, 'physical', 3, NULL, NULL, 2),
    ('Чистый код', 'Принципы и практики написания качественного кода', '978-5-04-123456-6', 1200.00, 'physical', 7, NULL, NULL, 3),
    ('Sapiens. Краткая история человечества', 'О том, как человек стал доминирующим видом на планете', '978-5-04-123456-7', 650.00, 'physical', 12, NULL, NULL, 3),
    ('Гарри Поттер и философский камень', 'Первая книга о мальчике, который выжил', '978-5-04-123456-8', 499.00, 'digital', NULL, 'books/harry-potter-1.epub', NULL, 1),
    ('Архитектура компьютера', 'Введение в организацию и архитектуру ЭВМ', '978-5-04-123456-9', 1500.00, 'physical', 4, NULL, NULL, 4),
    ('Собачье сердце', 'Повесть о профессоре Преображенском и Шарикове', '978-5-04-123456-0', 299.00, 'digital', NULL, 'books/heart-of-a-dog.epub', NULL, 1);

-- Book-Author links
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Преступление и наказание' AND a.name = 'Фёдор Достоевский';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Война и мир. Том 1' AND a.name = 'Лев Толстой';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Мастер и Маргарита' AND a.name = 'Михаил Булгаков';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Пикник на обочине' AND a.name = 'Братья Стругацкие';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Оно' AND a.name = 'Стивен Кинг';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Чистый код' AND a.name = 'Роберт Мартин';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Sapiens. Краткая история человечества' AND a.name = 'Юваль Ной Харари';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Гарри Поттер и философский камень' AND a.name = 'Джоан Роулинг';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Архитектура компьютера' AND a.name = 'Роберт Мартин';
INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id FROM bookshop.books b, bookshop.authors a WHERE b.title = 'Собачье сердце' AND a.name = 'Михаил Булгаков';

-- Book-Category links
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Преступление и наказание' AND c.slug = 'fiction';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Война и мир. Том 1' AND c.slug = 'fiction';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Война и мир. Том 1' AND c.slug = 'history';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Мастер и Маргарита' AND c.slug = 'fiction';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Пикник на обочине' AND c.slug = 'sci-fi';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Оно' AND c.slug = 'fiction';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Чистый код' AND c.slug = 'tech';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Sapiens. Краткая история человечества' AND c.slug = 'non-fiction';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Гарри Поттер и философский камень' AND c.slug = 'children';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Архитектура компьютера' AND c.slug = 'tech';
INSERT INTO bookshop.book_categories (book_id, category_id)
SELECT b.id, c.id FROM bookshop.books b, bookshop.categories c WHERE b.title = 'Собачье сердце' AND c.slug = 'fiction';

-- Customer user for testing (password: "customer123")
INSERT INTO bookshop.users (email, password_hash, full_name, phone_number, role)
VALUES (
    'customer@test.ru',
    '$2a$12$9Ea8K.kXo7.3aHb2O5CtReBN4Xmy/G7TkeLTK8LP9mP1YZWuc7R4q',
    'Иванов Иван Иванович',
    '+79001234567',
    'customer'
);

-- Address for customer
INSERT INTO bookshop.addresses (user_id, street_address, city, postal_code, country, is_default)
SELECT id, 'ул. Ленина, д. 10, кв. 5', 'Москва', '101000', 'Россия', TRUE
FROM bookshop.users WHERE email = 'customer@test.ru';

-- Reviews from customer
INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
SELECT b.id, u.id, 5, 'Шедевр на все времена', 'Очень глубокая книга, перечитывал несколько раз. Каждый раз открываю что-то новое.'
FROM bookshop.books b, bookshop.users u WHERE b.title = 'Преступление и наказание' AND u.email = 'customer@test.ru';

INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
SELECT b.id, u.id, 4, 'Интересно, но затянуто', 'Хорошая книга, но некоторые главы можно было сократить. В остальном — отлично.'
FROM bookshop.books b, bookshop.users u WHERE b.title = 'Война и мир. Том 1' AND u.email = 'customer@test.ru';

INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
SELECT b.id, u.id, 3, 'Странное впечатление', 'Неоднозначная книга. Местами гениально, местами непонятно. Но прочитать стоит.'
FROM bookshop.books b, bookshop.users u WHERE b.title = 'Мастер и Маргарита' AND u.email = 'customer@test.ru';

INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
SELECT b.id, u.id, 2, 'Не моё', 'Ожидал большего. Возможно, просто не мой жанр.'
FROM bookshop.books b, bookshop.users u WHERE b.title = 'Пикник на обочине' AND u.email = 'customer@test.ru';

INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
SELECT b.id, u.id, 5, 'Лучшая книга по программированию', 'Обязательна к прочтению каждым разработчиком. Меняет подход к написанию кода.'
FROM bookshop.books b, bookshop.users u WHERE b.title = 'Чистый код' AND u.email = 'customer@test.ru';
