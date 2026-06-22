BEGIN;

DROP INDEX IF EXISTS bookshop.idx_books_type_deleted;
DROP INDEX IF EXISTS bookshop.idx_books_price;
DROP INDEX IF EXISTS bookshop.idx_orders_created_at;
DROP INDEX IF EXISTS bookshop.idx_orders_user_id;
DROP INDEX IF EXISTS bookshop.idx_cart_items_user_id;

COMMIT;
