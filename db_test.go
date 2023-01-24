package sqltypes

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func connectDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS test_types (
			name TEXT NOT NULL PRIMARY KEY,
			value_str TEXT,
			value_int INTEGER
		)
	`)
	require.NoError(t, err)

	return db
}
