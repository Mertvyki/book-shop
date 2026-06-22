ALTER TABLE bookshop.books
    DROP CONSTRAINT fk_book_type,
    ADD COLUMN author VARCHAR(100);

UPDATE bookshop.books b
SET author = (
    SELECT a.name
    FROM bookshop.book_authors ba
    JOIN bookshop.authors a ON a.id = ba.author_id
    WHERE ba.book_id = b.id
    LIMIT 1
);

ALTER TABLE bookshop.books
    ALTER COLUMN author SET NOT NULL,
    DROP COLUMN publisher_id;

DROP TABLE bookshop.book_categories CASCADE;
DROP TABLE bookshop.categories CASCADE;
DROP TABLE bookshop.book_authors CASCADE;
DROP TABLE bookshop.authors CASCADE;
DROP TABLE bookshop.publishers CASCADE;
DROP TABLE bookshop.book_types CASCADE;
