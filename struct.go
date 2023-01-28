package sqltypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Struct[T any] struct {
	val *T
}

func NewStruct[T any](val *T) Struct[T] {
	return Struct[T]{val: val}
}

func (s *Struct[T]) Scan(value interface{}) error {
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
	*s = NewStruct(&val)

	return nil
}

func (s Struct[T]) Value() (driver.Value, error) {
	if s.val == nil {
		return nil, nil
	}
	return json.Marshal(s.val)
}

func (s *Struct[T]) Get() *T {
	return s.val
}

func (s *Struct[T]) Set(val *T) {
	s.val = val
}
