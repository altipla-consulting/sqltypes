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

	var st JSON[testJSON]
	st.V = &testJSON{Foo: "bar"}
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, `{"Foo":"bar"}`, val)
}

func TestJSONValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st JSON[testJSON]
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val sql.NullString
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.False(t, val.Valid)
}

func TestJSONScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", `{"Foo":"bar"}`)
	require.NoError(t, err)

	var st JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Foo, "bar")
}

func TestJSONScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var st JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.Nil(t, st.V)
}

func TestJSONSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st JSON[testJSON]
	st.V = &testJSON{Foo: "bar"}
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var other JSON[testJSON]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Foo, "bar")
}

func TestJSONArray(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	st := JSONArray[[]string]{V: []string{"foo", "bar"}}
	_, err := db.Exec("INSERT INTO test_types (name, value_str) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var other JSON[[]string]
	require.NoError(t, db.QueryRow("SELECT value_str FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.NotNil(t, st.V)
	require.Len(t, st.V, 2)
}
