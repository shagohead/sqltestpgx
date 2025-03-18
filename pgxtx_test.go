package sqltestpgx

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestTxQuery(t *testing.T) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		conn.Close(ctx)
	})
	pgxtx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pgxtx.Rollback(ctx)
	})
	tx := Tx(pgxtx)

	for _, test := range []struct {
		name  string // Test name
		query string // Input query
		args  []any  // Input arguments
		want  string // Expected strings output
	}{
		{
			name:  "one-row/multiple-types",
			query: `SELECT 1, 'test', NULL`,
			want:  `[1 test <nil>]`,
		},
		{
			name:  "multiple-rows",
			query: "SELECT t FROM generate_series(1, 5) AS t",
			want:  "[1] [2] [3] [4] [5]",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			rows, err := tx.Query(ctx, test.query, test.args...)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			var res []string
			for rows.Next() {
				v, err := rows.String()
				if err != nil {
					t.Fatal(err)
				}
				res = append(res, v)
			}
			if got := strings.Join(res, " "); got != test.want {
				t.Errorf("rows String() output: %q, want: %q", got, test.want)
			}
		})
	}
}
