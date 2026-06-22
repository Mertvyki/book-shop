package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
	"golang.org/x/sync/errgroup"
)

func uniqueIDs(ids []int) []int {
	seen := make(map[int]struct{}, len(ids))
	res := make([]int, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			res = append(res, id)
		}
	}
	return res
}

func (r *BooksRepository) getAuthorsByBookIDs(ctx context.Context, bookIDs []int) (map[int][]domain.Author, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT ba.book_id, a.id, a.name, a.bio, a.birth_year, a.created_at
		FROM bookshop.authors a
		JOIN bookshop.book_authors ba ON ba.author_id = a.id
		WHERE ba.book_id = ANY($1)
		ORDER BY a.name
	`, bookIDs)
	if err != nil {
		return nil, fmt.Errorf("batch get authors: %w", err)
	}
	defer rows.Close()

	result := make(map[int][]domain.Author)
	for rows.Next() {
		var bookID int
		var m AuthorModel
		if err := rows.Scan(&bookID, &m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt); err == nil {
			result[bookID] = append(result[bookID], m.ToDomain())
		}
	}
	return result, nil
}

func (r *BooksRepository) getCategoriesByBookIDs(ctx context.Context, bookIDs []int) (map[int][]domain.Category, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT bc.book_id, c.id, c.name, c.slug, c.description
		FROM bookshop.categories c
		JOIN bookshop.book_categories bc ON bc.category_id = c.id
		WHERE bc.book_id = ANY($1)
		ORDER BY c.name
	`, bookIDs)
	if err != nil {
		return nil, fmt.Errorf("batch get categories: %w", err)
	}
	defer rows.Close()

	result := make(map[int][]domain.Category)
	for rows.Next() {
		var bookID int
		var m CategoryModel
		if err := rows.Scan(&bookID, &m.ID, &m.Name, &m.Slug, &m.Description); err == nil {
			result[bookID] = append(result[bookID], m.ToDomain())
		}
	}
	return result, nil
}

func (r *BooksRepository) getPublishersByIDs(ctx context.Context, publisherIDs []int) (map[int]domain.Publisher, error) {
	if len(publisherIDs) == 0 {
		return nil, nil
	}
	rows, err := r.pool.Query(ctx, `
		SELECT id, name FROM bookshop.publishers WHERE id = ANY($1)
	`, publisherIDs)
	if err != nil {
		return nil, fmt.Errorf("batch get publishers: %w", err)
	}
	defer rows.Close()

	result := make(map[int]domain.Publisher, len(publisherIDs))
	for rows.Next() {
		var m PublisherModel
		if err := rows.Scan(&m.ID, &m.Name); err == nil {
			result[m.ID] = m.ToDomain()
		}
	}
	return result, nil
}

type booksQueryParts struct {
	from   string
	where  []string
	args   []any
	argPos int
}

func (r *BooksRepository) buildBooksQuery(qp books_service.GetBooksQueryParams) booksQueryParts {
	args := []any{}
	argPos := 1

	fromClause := `FROM bookshop.books b`

	whereClauses := []string{"b.deleted_at IS NULL"}

	if qp.AuthorID != nil {
		fromClause += ` JOIN bookshop.book_authors ba_author ON ba_author.book_id = b.id`
		whereClauses = append(whereClauses, fmt.Sprintf("ba_author.author_id = $%d", argPos))
		args = append(args, *qp.AuthorID)
		argPos++
	}

	if qp.CategoryID != nil {
		fromClause += ` JOIN bookshop.book_categories bc_cat ON bc_cat.book_id = b.id`
		whereClauses = append(whereClauses, fmt.Sprintf("bc_cat.category_id = $%d", argPos))
		args = append(args, *qp.CategoryID)
		argPos++
	}

	if qp.PublisherID != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("b.publisher_id = $%d", argPos))
		args = append(args, *qp.PublisherID)
		argPos++
	}

	if qp.Type != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("b.book_type = $%d", argPos))
		args = append(args, *qp.Type)
		argPos++
	}

	if qp.Search != nil {
		whereClauses = append(whereClauses, fmt.Sprintf(
			"(b.title ILIKE $%d OR EXISTS (SELECT 1 FROM bookshop.book_authors ba_s JOIN bookshop.authors a_s ON a_s.id = ba_s.author_id WHERE ba_s.book_id = b.id AND a_s.name ILIKE $%d))",
			argPos, argPos,
		))
		args = append(args, "%"+*qp.Search+"%")
		argPos++
	}

	if qp.MinPrice != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("b.price >= $%d", argPos))
		args = append(args, *qp.MinPrice)
		argPos++
	}

	if qp.MaxPrice != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("b.price <= $%d", argPos))
		args = append(args, *qp.MaxPrice)
		argPos++
	}

	return booksQueryParts{
		from:   fromClause,
		where:  whereClauses,
		args:   args,
		argPos: argPos,
	}
}

func (r *BooksRepository) CountBooks(ctx context.Context, qp books_service.GetBooksQueryParams) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	q := r.buildBooksQuery(qp)

	query := fmt.Sprintf(`SELECT COUNT(DISTINCT b.id) %s WHERE %s`,
		q.from, strings.Join(q.where, " AND "))

	var total int
	err := r.pool.QueryRow(ctx, query, q.args...).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count books: %w", err)
	}

	return total, nil
}

func (r *BooksRepository) GetBooks(
	ctx context.Context,
	qp books_service.GetBooksQueryParams,
) ([]domain.Book, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	q := r.buildBooksQuery(qp)

	offset := (qp.Page - 1) * qp.Limit

	orderBy := "b.created_at DESC"
	bestSellerExpr := ""
	if qp.Sort != nil {
		switch *qp.Sort {
		case "price_asc":
			orderBy = "b.price ASC"
		case "price_desc":
			orderBy = "b.price DESC"
		case "newest":
			orderBy = "b.created_at DESC"
		case "bestsellers":
			bestSellerExpr = `, COALESCE((SELECT COUNT(oi.id) FROM bookshop.order_items oi WHERE oi.book_id = b.id), 0) AS sale_count`
			orderBy = "sale_count DESC"
		}
	}

	query := fmt.Sprintf(`
	SELECT DISTINCT b.id, b.version, b.title, b.description, b.isbn, b.price,
		b.book_type, b.stock_quantity, b.file_key, b.cover_image_key, b.publisher_id,
		COALESCE(rv.avg_rating, 0),
		b.created_at,
		COUNT(*) OVER() AS total_count
		%s
	%s
	LEFT JOIN LATERAL (SELECT AVG(r.rating) AS avg_rating FROM bookshop.reviews r WHERE r.book_id = b.id) rv ON true
	WHERE %s
	ORDER BY %s
	LIMIT $%d OFFSET $%d
	`, bestSellerExpr, q.from, strings.Join(q.where, " AND "), orderBy, q.argPos, q.argPos+1)

	args := append(q.args, qp.Limit, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("query books: %w", err)
	}
	defer rows.Close()

	books := make([]domain.Book, 0)
	bookIDs := make([]int, 0)
	var totalCount int
	var saleCount int

	for rows.Next() {
		var model BookModel
		if bestSellerExpr != "" {
			err = rows.Scan(
				&model.ID, &model.Version, &model.Title, &model.Description,
				&model.ISBN, &model.Price, &model.BookType, &model.StockQuantity,
				&model.FileURL, &model.CoverImageURL, &model.PublisherID,
				&model.AvgRating, &model.CreatedAt,
				&totalCount, &saleCount,
			)
		} else {
			err = rows.Scan(
				&model.ID, &model.Version, &model.Title, &model.Description,
				&model.ISBN, &model.Price, &model.BookType, &model.StockQuantity,
				&model.FileURL, &model.CoverImageURL, &model.PublisherID,
				&model.AvgRating, &model.CreatedAt,
				&totalCount,
			)
		}
		if err != nil {
			return nil, 0, fmt.Errorf("scan book: %w", err)
		}

		books = append(books, model.ToDomain())
		bookIDs = append(bookIDs, model.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows iteration: %w", err)
	}

	// Batch fetch related data concurrently (avoids N+1 queries)
	if len(bookIDs) > 0 {
		var (
			authorsMap    map[int][]domain.Author
			categoriesMap map[int][]domain.Category
			publishersMap map[int]domain.Publisher
		)

		g, gctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			var err error
			authorsMap, err = r.getAuthorsByBookIDs(gctx, bookIDs)
			return err
		})

		g.Go(func() error {
			var err error
			categoriesMap, err = r.getCategoriesByBookIDs(gctx, bookIDs)
			return err
		})

		pubIDs := make([]int, 0, len(books))
		for _, b := range books {
			if b.PublisherID != nil {
				pubIDs = append(pubIDs, *b.PublisherID)
			}
		}
		pubIDs = uniqueIDs(pubIDs)

		if len(pubIDs) > 0 {
			g.Go(func() error {
				var err error
				publishersMap, err = r.getPublishersByIDs(gctx, pubIDs)
				return err
			})
		}

		g.Wait()

		if authorsMap != nil {
			for i := range books {
				if a, ok := authorsMap[books[i].ID]; ok {
					books[i].Authors = a
				}
			}
		}

		if categoriesMap != nil {
			for i := range books {
				if c, ok := categoriesMap[books[i].ID]; ok {
					books[i].Categories = c
				}
			}
		}

		if publishersMap != nil {
			for i := range books {
				if books[i].PublisherID != nil {
					if p, ok := publishersMap[*books[i].PublisherID]; ok {
						books[i].Publisher = &p
					}
				}
			}
		}
	}

	return books, totalCount, nil
}
