CREATE TABLE bookshop.reviews (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    book_id INT NOT NULL REFERENCES bookshop.books(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES bookshop.users(id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    title TEXT,
    body TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(book_id, user_id)
);
