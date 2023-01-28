package sqltypes

import (
	"database/sql/driver"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)

type Timestamp time.Time

func (t *Timestamp) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value := value.(type) {
	case int64:
		read := time.Unix(value, 0)
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

type Date civil.Date

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value := value.(type) {
	case string:
		read, err := civil.ParseDate(value)
		if err != nil {
			return fmt.Errorf("sqltypes: cannot parse %q: %w", value, err)
		}
		*d = Date(read)
		return nil
	default:
		return fmt.Errorf("sqltypes: unknown time type %T", value)
	}
}

func (d Date) Value() (driver.Value, error) {
	if d.Date().IsZero() {
		return "", nil
	}
	return d.String(), nil
}

func (d Date) Date() civil.Date {
	return civil.Date(d)
}

func (d Date) String() string {
	return d.Date().String()
}
