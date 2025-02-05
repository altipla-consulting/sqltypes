package sqltypes

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Proto[M proto.Message] struct {
	V M
}

func NewProto[M proto.Message](val M) Proto[M] {
	return Proto[M]{V: val}
}

func (s *Proto[M]) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	s.V = s.V.ProtoReflect().New().Interface().(M)

	opts := protojson.UnmarshalOptions{DiscardUnknown: true}
	switch v := value.(type) {
	case string:
		s.V = proto.Clone(s.V).(M)
		return opts.Unmarshal([]byte(v), s.V)
	case []byte:
		s.V = proto.Clone(s.V).(M)
		return opts.Unmarshal(v, s.V)
	}

	return fmt.Errorf("unsupported protobuf column type %T", value)
}

func (s Proto[M]) Value() (driver.Value, error) {
	if proto.Equal(s.V.ProtoReflect().Type().Zero().Interface().(proto.Message), s.V) {
		return nil, nil
	}

	b, err := protojson.Marshal(s.V)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
