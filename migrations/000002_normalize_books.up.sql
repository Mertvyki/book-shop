CREATE TABLE bookshop.book_types (
    code VARCHAR(20) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

INSERT INTO bookshop.book_types (code, name) VALUES
    ('digital', 'Digital'),
    ('physical', 'Physical');

CREATE TABLE bookshop.publishers (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL
);

CREATE TABLE bookshop.authors (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    bio        TEXT,
    birth_year SMALLINT
);

CREATE TABLE bookshop.book_authors (
    book_id   INTEGER NOT NULL REFERENCES bookshop.books(id) ON DELETE CASCADE,
    author_id INTEGER NOT NULL REFERENCES bookshop.authors(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

CREATE TABLE bookshop.categories (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL UNIQUE,
    slug        VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE bookshop.book_categories (
    book_id     INTEGER NOT NULL REFERENCES bookshop.books(id) ON DELETE CASCADE,
    category_id INTEGER NOT NULL REFERENCES bookshop.categories(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, category_id)
);

INSERT INTO bookshop.authors (name)
SELECT DISTINCT author FROM bookshop.books WHERE author IS NOT NULL;

INSERT INTO bookshop.book_authors (book_id, author_id)
SELECT b.id, a.id
FROM bookshop.books b
JOIN bookshop.authors a ON a.name = b.author;

ALTER TABLE bookshop.books
    ADD COLUMN publisher_id INTEGER REFERENCES bookshop.publishers(id),
    DROP COLUMN author,
    ADD CONSTRAINT fk_book_type FOREIGN KEY (book_type) REFERENCES bookshop.book_types(code);
