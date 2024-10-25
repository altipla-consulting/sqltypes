package sqltypes

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Foo string
}

func TestStructValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st Struct[testStruct]
	st.Set(&testStruct{Foo: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_blob) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_blob FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, `{"Foo":"bar"}`, val)
}

func TestStructValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st Struct[testStruct]
	_, err := db.Exec("INSERT INTO test_types (name, value_blob) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val sql.NullString
	require.NoError(t, db.QueryRow("SELECT value_blob FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.False(t, val.Valid)
}

func TestStructScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_blob) VALUES (?, ?)", "foo", []byte(`{"Foo":"bar"}`))
	require.NoError(t, err)

	var st Struct[testStruct]
	require.NoError(t, db.QueryRow("SELECT value_blob FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.NotNil(t, st.Get())
	require.Equal(t, st.Get().Foo, "bar")
}

func TestStructScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_blob) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var st Struct[testStruct]
	require.NoError(t, db.QueryRow("SELECT value_blob FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.Nil(t, st.Get())
}

func TestStructSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st Struct[testStruct]
	st.Set(&testStruct{Foo: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_blob) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var other Struct[testStruct]
	require.NoError(t, db.QueryRow("SELECT value_blob FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.NotNil(t, st.Get())
	require.Equal(t, st.Get().Foo, "bar")
}
