package sqltypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON[T any] struct {
	V *T
}

func NewJSON[T any](val *T) JSON[T] {
	return JSON[T]{V: val}
}

func (s *JSON[T]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var val T
	switch value.(type) {
	case []byte:
		if err := json.Unmarshal(value.([]byte), &val); err != nil {
			return fmt.Errorf("sqltypes: cannot unmarshal struct: %w", err)
		}
	case string:
		if err := json.Unmarshal([]byte(value.(string)), &val); err != nil {
			return fmt.Errorf("sqltypes: cannot unmarshal struct: %w", err)
		}
	default:
		return fmt.Errorf("sqltypes: unknown array type %T", value)
	}
	*s = NewJSON(&val)

	return nil
}

func (s JSON[T]) Value() (driver.Value, error) {
	if s.V == nil {
		return nil, nil
	}
	value, err := json.Marshal(s.V)
	if err != nil {
		return nil, fmt.Errorf("sqltypes: cannot marshal type: %w", err)
	}
	return string(value), nil
}

type JSONArray[T any] []T

func (arr *JSONArray[T]) Scan(value interface{}) error {
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

func (arr JSONArray[T]) Value() (driver.Value, error) {
	if arr == nil {
		return "[]", nil
	}
	v, err := json.Marshal(arr)
	if err != nil {
		return nil, fmt.Errorf("sqltypes: cannot marshal array: %w", err)
	}
	return string(v), nil
}

func (arr JSONArray[T]) Unwrap() []T {
	return arr
}
