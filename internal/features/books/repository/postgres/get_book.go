package book_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) GetBook(ctx context.Context, id int) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var model BookModel
	err := r.pool.QueryRow(ctx, `
		SELECT b.id, b.version, b.title, b.description, b.isbn, b.price, b.book_type,
			b.stock_quantity, b.file_key, b.cover_image_key, b.publisher_id,
			COALESCE((SELECT AVG(r.rating) FROM bookshop.reviews r WHERE r.book_id = b.id), 0),
			b.created_at
		FROM bookshop.books b
		WHERE b.id = $1 AND b.deleted_at IS NULL
	`, id).Scan(
		&model.ID, &model.Version, &model.Title, &model.Description,
		&model.ISBN, &model.Price, &model.BookType, &model.StockQuantity,
		&model.FileURL, &model.CoverImageURL, &model.PublisherID,
		&model.AvgRating, &model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Book{}, fmt.Errorf("book with id=%d: %w", id, core_errors.ErrNotFound)
		}

		return domain.Book{}, fmt.Errorf("scan book: %w", err)
	}

	book := model.ToDomain()

	authorsMap, err := r.getAuthorsByBookIDs(ctx, []int{id})
	if err == nil {
		if a, ok := authorsMap[id]; ok {
			book.Authors = a
		}
	}

	categoriesMap, err := r.getCategoriesByBookIDs(ctx, []int{id})
	if err == nil {
		if c, ok := categoriesMap[id]; ok {
			book.Categories = c
		}
	}

	if model.PublisherID != nil {
		publishersMap, err := r.getPublishersByIDs(ctx, []int{*model.PublisherID})
		if err == nil {
			if p, ok := publishersMap[*model.PublisherID]; ok {
				book.Publisher = &p
			}
		}
	}

	return book, nil
}

func (r *BooksRepository) getBookAuthors(ctx context.Context, bookID int) []domain.Author {
	rows, err := r.pool.Query(ctx, `
		SELECT a.id, a.name, a.bio, a.birth_year, a.created_at
		FROM bookshop.authors a
		JOIN bookshop.book_authors ba ON ba.author_id = a.id
		WHERE ba.book_id = $1
		ORDER BY a.name
	`, bookID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	authors := make([]domain.Author, 0)
	for rows.Next() {
		var m AuthorModel
		if err := rows.Scan(&m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt); err == nil {
			authors = append(authors, m.ToDomain())
		}
	}

	return authors
}

func (r *BooksRepository) getBookCategories(ctx context.Context, bookID int) []domain.Category {
	rows, err := r.pool.Query(ctx, `
		SELECT c.id, c.name, c.slug, c.description
		FROM bookshop.categories c
		JOIN bookshop.book_categories bc ON bc.category_id = c.id
		WHERE bc.book_id = $1
		ORDER BY c.name
	`, bookID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	categories := make([]domain.Category, 0)
	for rows.Next() {
		var m CategoryModel
		if err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Description); err == nil {
			categories = append(categories, m.ToDomain())
		}
	}

	return categories
}

func (r *BooksRepository) GetPublisher(ctx context.Context, id int) (domain.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m PublisherModel
	err := r.pool.QueryRow(ctx, `SELECT id, name FROM bookshop.publishers WHERE id = $1`, id).
		Scan(&m.ID, &m.Name)
	if err != nil {
		return domain.Publisher{}, fmt.Errorf("get publisher: %w", err)
	}

	return m.ToDomain(), nil
}
