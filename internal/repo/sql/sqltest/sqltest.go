package sqltest

import (
	"database/sql"
	"fmt"
	"testing"

	// 3rd party
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/stretchr/testify/require"
)

// DB struct represents a wrapper for a db handle
type DB struct {
	Db *sql.DB
}

func (tdb *DB) RunMigrations(source string) error {
	driver, err := postgres.WithInstance(tdb.Db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		source,
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}

// RequireTotalRows counts the rows in given table
func (tdb *DB) RequireTotalRows(t *testing.T, table string, expectedCount int) {
	var n int
	err := tdb.Db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&n)
	require.NoError(t, err)
	require.Equal(t, expectedCount, n)
}
