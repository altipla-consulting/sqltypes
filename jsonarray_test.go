package sqltypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
