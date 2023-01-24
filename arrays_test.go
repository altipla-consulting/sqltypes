package sqltypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var arr Array[int64] = []int64{1, 2, 3}
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "[1,2,3]", val)
}

func TestArrayValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var arr Array[int64]
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "[]", val)
}

func TestArrayScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", "[1, 2, 3]")
	require.NoError(t, err)

	var arr Array[int64]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&arr))

	require.Len(t, arr, 3)
	require.EqualValues(t, arr[0], 1)
	require.EqualValues(t, arr[1], 2)
	require.EqualValues(t, arr[2], 3)
}

func TestArrayScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var arr Array[int64]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&arr))

	require.Len(t, arr, 0)
}

func TestArraySaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var arr Array[int64] = []int64{1, 2, 3}
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var other Array[int64]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.Len(t, other, 3)
	require.EqualValues(t, other[0], 1)
	require.EqualValues(t, other[1], 2)
	require.EqualValues(t, other[2], 3)
}
