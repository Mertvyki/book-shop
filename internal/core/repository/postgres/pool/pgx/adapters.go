package core_pgx_pool

import (
	"context"
	"errors"
	"fmt"

	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

type pgxTx struct {
	pgx.Tx
}

func (t pgxTx) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	row := t.Tx.QueryRow(ctx, sql, args...)
	return pgxRow{row}
}

func (t pgxTx) Exec(ctx context.Context, sql string, arguments ...any) (core_postgres_pool.CommandTag, error) {
	tag, err := t.Tx.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}

	return pgxCommandTag{tag}, nil
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKeyErrorCode = "23503"
		pgxUniqueViolationErrorCode    = "23505"
		pgxCheckViolationErrorCode     = "23514"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgxViolatesForeignKeyErrorCode:
			return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrViolatesForeignKey)
		case pgxUniqueViolationErrorCode:
			return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrUniqueViolation)
		case pgxCheckViolationErrorCode:
			return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrCheckViolation)
		}
	}

	return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrUnknown)
}
