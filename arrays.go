package sqltypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Depcreated: use JSONArray instead
type Array[T any] []T

func (arr *Array[T]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), &arr)
	case string:
		return json.Unmarshal([]byte(value.(string)), &arr)
	default:
		return fmt.Errorf("sqltypes: unknown array type %T", value)
	}
}

func (arr Array[T]) Value() (driver.Value, error) {
	if arr == nil {
		return "[]", nil
	}
	return json.Marshal(arr)
}

func (arr Array[T]) Unwrap() []T {
	return arr
}
