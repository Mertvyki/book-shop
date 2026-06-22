BEGIN;

CREATE INDEX IF NOT EXISTS idx_books_type_deleted
    ON bookshop.books (book_type, deleted_at);

CREATE INDEX IF NOT EXISTS idx_books_price
    ON bookshop.reviews (book_id, user_id);

CREATE INDEX IF NOT EXISTS idx_orders_created_at
    ON bookshop.orders (created_at DESC);

CREATE INDEX IF NOT EXISTS idx_orders_user_id
    ON bookshop.orders (user_id);

CREATE INDEX IF NOT EXISTS idx_cart_items_user_id
    ON bookshop.cart_items (user_id);

COMMIT;
