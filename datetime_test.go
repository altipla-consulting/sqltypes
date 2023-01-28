package sqltypes

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"
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

func TestDateValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var date = Date(civil.Date{Year: 2018, Month: 1, Day: 2})
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", date)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "2018-01-02", val)
}

func TestDateValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var date Date
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", date)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "", val)
}

func TestDateScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", "2018-01-02")
	require.NoError(t, err)

	var date Date
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&date))

	require.Equal(t, date.Date().Year, 2018)
	require.Equal(t, date.Date().Month, time.January)
	require.Equal(t, date.Date().Day, 2)
}

func TestDateScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var date Date
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&date))

	require.True(t, date.Date().IsZero())
}

func TestDateSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var date = Date(civil.Date{Year: 2018, Month: 1, Day: 2})
	_, err := db.Exec("INSERT INTO test_types (name, value_int) VALUES (?, ?)", "foo", date)
	require.NoError(t, err)

	var other Date
	require.NoError(t, db.QueryRow("SELECT value_int FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.Equal(t, date.Date().Year, 2018)
	require.Equal(t, date.Date().Month, time.January)
	require.Equal(t, date.Date().Day, 2)
}
