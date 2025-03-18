sqltest with pgx driver
=======================

Implementation of sqltest.Tx/Rows for pgx/v5 database driver.


Usage example:

```go
import (
	"testing"

	"github.com/shagohead/sqltest"
	"github.com/shagohead/sqltestpgx"
)

// TestSchema tests database schema with queries from testdata/*.sql
func TestSchema(t *testing.T) {
	set, err := sqltest.DefaultFileSet()
	if err != nil {
		t.Fatal(err)
	}
	for name, test := range set.All() {
		t.Run(name, func(t *testing.T) {
			// dbtest.StartTx is a helper which creates and rollbacks transactions for tests.
			err := test.Run(sqltestpgx.Tx(dbtest.StartTx(t)))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
```
