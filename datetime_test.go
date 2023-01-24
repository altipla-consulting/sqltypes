package sqltypes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimestampValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var ts = Timestamp(time.Date(2018, 1, 2, 3, 4, 5, 6, time.UTC))
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", ts)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "1514862245", val)
}

func TestTimestampValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var ts Timestamp
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", ts)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "0", val)
}

func TestTimestampScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", 1514862245)
	require.NoError(t, err)

	var ts Timestamp
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&ts))

	require.WithinDuration(t, ts.Time(), time.Date(2018, 1, 2, 3, 4, 5, 0, time.UTC), time.Second)
}

func TestTimestampScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var ts Timestamp
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&ts))

	require.True(t, ts.Time().IsZero())
}

func TestTimestampSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var ts = Timestamp(time.Date(2018, 1, 2, 3, 4, 5, 6, time.UTC))
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", ts)
	require.NoError(t, err)

	var other Timestamp
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.WithinDuration(t, other.Time(), time.Date(2018, 1, 2, 3, 4, 5, 0, time.UTC), time.Second)
}
