CREATE SCHEMA bookshop;

CREATE TABLE bookshop.users (
    id              SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    full_name       VARCHAR(50)  NOT NULL,
    phone_number    VARCHAR(20),
    role            VARCHAR(20)  NOT NULL DEFAULT 'customer',
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_full_name_length CHECK (char_length(full_name) BETWEEN 15 AND 50),

    CONSTRAINT chk_email_format CHECK (
        email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
    ),

    CONSTRAINT chk_phone_format CHECK (
        phone_number IS NULL OR phone_number ~ '^\+?[0-9]{7,15}$'
    ),

    CONSTRAINT chk_user_role CHECK (role IN ('customer', 'admin'))
);

CREATE TABLE bookshop.books (
    id               SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title            VARCHAR(200) NOT NULL,
    author           VARCHAR(100) NOT NULL,
    description      TEXT,
    isbn             VARCHAR(20) UNIQUE,
    price            DECIMAL(10,2) NOT NULL,
    book_type        VARCHAR(10)  NOT NULL,   -- 'digital' или 'physical'
    stock_quantity   INTEGER,                -- остаток (только для physical)
    file_url         TEXT,                   -- ссылка на файл (для digital)
    cover_image_url  TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_price_nonnegative CHECK (price >= 0),

    CONSTRAINT chk_book_type CHECK (book_type IN ('digital', 'physical')),

    CONSTRAINT chk_stock_quantity CHECK (
        (book_type = 'physical' AND stock_quantity IS NOT NULL AND stock_quantity >= 0)
        OR (book_type = 'digital' AND stock_quantity IS NULL)
    )
);

CREATE TABLE bookshop.addresses (
    id             SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    user_id        INTEGER      NOT NULL REFERENCES bookshop.users(id) ON DELETE CASCADE,
    street_address VARCHAR(255) NOT NULL,
    city           VARCHAR(100) NOT NULL,
    postal_code    VARCHAR(20)  NOT NULL,
    country        VARCHAR(100) NOT NULL DEFAULT 'Россия',
    is_default     BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_addresses_user_id ON bookshop.addresses(user_id);

CREATE TABLE bookshop.orders (
    id                  SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    user_id             INTEGER        NOT NULL REFERENCES bookshop.users(id),
    status              VARCHAR(20)    NOT NULL DEFAULT 'pending',
    total_amount        DECIMAL(10,2)  NOT NULL,
    shipping_address_id INTEGER        REFERENCES bookshop.addresses(id) ON DELETE SET NULL,
    payment_method      VARCHAR(50),
    created_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_total_amount_nonnegative CHECK (total_amount >= 0),

    CONSTRAINT chk_order_status CHECK (
        status IN ('pending', 'paid', 'shipped', 'delivered', 'cancelled')
    )
);

CREATE INDEX idx_orders_user_id ON bookshop.orders(user_id);
CREATE INDEX idx_orders_status ON bookshop.orders(status);

CREATE TABLE bookshop.order_items (
    id         SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    order_id   INTEGER        NOT NULL REFERENCES bookshop.orders(id) ON DELETE CASCADE,
    book_id    INTEGER        NOT NULL REFERENCES bookshop.books(id),
    quantity   INTEGER        NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    item_type  VARCHAR(10)   NOT NULL,

    CONSTRAINT chk_quantity_positive CHECK (quantity > 0),

    CONSTRAINT chk_unit_price_nonnegative CHECK (unit_price >= 0),

    CONSTRAINT chk_item_type CHECK (item_type IN ('digital', 'physical'))
);

CREATE INDEX idx_order_items_order_id ON bookshop.order_items(order_id);
CREATE INDEX idx_order_items_book_id ON bookshop.order_items(book_id);

CREATE TABLE bookshop.cart_items (
    id         SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    user_id    INTEGER        NOT NULL REFERENCES bookshop.users(id) ON DELETE CASCADE,
    book_id    INTEGER        NOT NULL REFERENCES bookshop.books(id),
    quantity   INTEGER        NOT NULL,
    added_at   TIMESTAMPTZ    NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, book_id),

    CONSTRAINT chk_cart_quantity CHECK (quantity > 0)
);

CREATE INDEX idx_cart_items_user_id ON bookshop.cart_items(user_id);