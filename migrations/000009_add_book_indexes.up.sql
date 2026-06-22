BEGIN;

CREATE INDEX IF NOT EXISTS idx_books_deleted_at
    ON bookshop.books (deleted_at);

CREATE INDEX IF NOT EXISTS idx_books_publisher_id
    ON bookshop.books (publisher_id);

CREATE INDEX IF NOT EXISTS idx_books_book_type
    ON bookshop.books (book_type);

CREATE INDEX IF NOT EXISTS idx_book_authors_book_id
    ON bookshop.book_authors (book_id);

CREATE INDEX IF NOT EXISTS idx_book_authors_author_id
    ON bookshop.book_authors (author_id);

CREATE INDEX IF NOT EXISTS idx_book_categories_book_id
    ON bookshop.book_categories (book_id);

CREATE INDEX IF NOT EXISTS idx_book_categories_category_id
    ON bookshop.book_categories (category_id);

CREATE INDEX IF NOT EXISTS idx_reviews_book_id
    ON bookshop.reviews (book_id);

COMMIT;
