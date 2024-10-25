package sqltypes

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

type testJSON struct {
	Foo string
}

func TestJSONValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	st := NewJSON(&testJSON{Foo: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, `{"Foo":"bar"}`, val)
}

func TestJSONValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st JSON[testJSON]
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val sql.NullString
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.False(t, val.Valid)
}

func TestJSONScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", `{"Foo":"bar"}`)
	require.NoError(t, err)

	var st JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Foo, "bar")
}

func TestJSONScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var st JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.Nil(t, st.V)
}

func TestJSONSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	st := NewJSON(&testJSON{Foo: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var other JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Foo, "bar")
}

func TestJSONArrayValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	arr := JSONArray[int64]{1, 2, 3}
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "[1,2,3]", val)
}

func TestJSONArrayValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var arr JSONArray[int64]
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, "[]", val)
}

func TestJSONArrayScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", "[1, 2, 3]")
	require.NoError(t, err)

	var arr JSONArray[int64]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&arr))

	require.Len(t, arr, 3)
	require.EqualValues(t, arr[0], 1)
	require.EqualValues(t, arr[1], 2)
	require.EqualValues(t, arr[2], 3)
}

func TestJSONArrayScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var arr JSONArray[int64]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&arr))

	require.Len(t, arr, 0)
}

func TestJSONArraySaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	arr := JSONArray[int64]{1, 2, 3}
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", arr)
	require.NoError(t, err)

	var other JSONArray[int64]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.Len(t, other, 3)
	require.EqualValues(t, other[0], 1)
	require.EqualValues(t, other[1], 2)
	require.EqualValues(t, other[2], 3)
}
