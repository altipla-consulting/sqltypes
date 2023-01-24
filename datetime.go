package sqltypes

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case int64:
		read := time.Unix(value.(int64), 0)
		*t = Timestamp(read)
		return nil
	default:
		return fmt.Errorf("sqltypes: unknown time type %T", value)
	}
}

func (t Timestamp) Value() (driver.Value, error) {
	if t.Time().IsZero() {
		return "0", nil
	}
	return t.Time().Unix(), nil
}

func (t Timestamp) Time() time.Time {
	return time.Time(t)
}
