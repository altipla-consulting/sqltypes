package sqltypes

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	pb "google.golang.org/protobuf/types/known/apipb"
)

func TestProtoValue(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	st := NewProto(&pb.Api{Name: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val string
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.Equal(t, `{"name":"bar"}`, val)
}

func TestProtoValueNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	var st Proto[*pb.Api]
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var val sql.NullString
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&val))

	require.False(t, val.Valid)
}

func TestProtoScan(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", `{"name":"bar"}`)
	require.NoError(t, err)

	var st Proto[*pb.Api]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Name, "bar")
}

func TestProtoScanNil(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", nil)
	require.NoError(t, err)

	var st Proto[*pb.Api]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&st))

	require.Nil(t, st.V)
}

func TestProtoSaveLoad(t *testing.T) {
	db := connectDB(t)
	defer db.Close()

	st := NewProto(&pb.Api{Name: "bar"})
	_, err := db.Exec("INSERT INTO test_types (name, value_text) VALUES (?, ?)", "foo", st)
	require.NoError(t, err)

	var other Proto[*pb.Api]
	require.NoError(t, db.QueryRow("SELECT value_text FROM test_types WHERE name = ?", "foo").Scan(&other))

	require.NotNil(t, st.V)
	require.Equal(t, st.V.Name, "bar")
}
