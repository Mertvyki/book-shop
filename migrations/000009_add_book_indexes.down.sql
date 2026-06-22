BEGIN;

DROP INDEX IF EXISTS bookshop.idx_books_deleted_at;
DROP INDEX IF EXISTS bookshop.idx_books_publisher_id;
DROP INDEX IF EXISTS bookshop.idx_books_book_type;
DROP INDEX IF EXISTS bookshop.idx_book_authors_book_id;
DROP INDEX IF EXISTS bookshop.idx_book_authors_author_id;
DROP INDEX IF EXISTS bookshop.idx_book_categories_book_id;
DROP INDEX IF EXISTS bookshop.idx_book_categories_category_id;
DROP INDEX IF EXISTS bookshop.idx_reviews_book_id;

COMMIT;
