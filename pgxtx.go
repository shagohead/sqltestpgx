package sqltestpgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/shagohead/sqltest"
)

func Tx(pgtx pgx.Tx) sqltest.Tx {
	return &tx{tx: pgtx}
}

type tx struct {
	tx pgx.Tx
}

// Exec implements sqltest.Tx.
func (t *tx) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := t.tx.Exec(ctx, sql, args...)
	return err
}

type rows struct {
	pgrows pgx.Rows
}

// Close implements sqltest.Rows.
func (r *rows) Close() {
	r.pgrows.Close()
}

// Err implements sqltest.Rows.
func (r *rows) Err() error {
	return r.pgrows.Err()
}

// Next implements sqltest.Rows.
func (r *rows) Next() bool {
	return r.pgrows.Next()
}

// String implements sqltest.Rows.
func (r *rows) String() (string, error) {
	v, err := r.pgrows.Values()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(v), nil
}

// Query implements sqltest.Tx.
func (t *tx) Query(ctx context.Context, sql string, args ...any) (sqltest.Rows, error) {
	r, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &rows{pgrows: r}, nil
}

var _ sqltest.Tx = (*tx)(nil)
