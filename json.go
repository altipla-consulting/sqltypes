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
		return nil, err
	}
	return string(value), nil
}

type JSONArray[T any] struct {
	V T
}

func NewJSONArray[T any](val *T) JSONArray[T] {
	return JSONArray[T]{V: *val}
}

func (s *JSONArray[T]) Scan(value interface{}) error {
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
	*s = NewJSONArray(&val)

	return nil
}

func (s JSONArray[T]) Value() (driver.Value, error) {
	value, err := json.Marshal(s.V)
	if err != nil {
		return nil, err
	}
	return string(value), nil
}
