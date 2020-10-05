package ptypes

import (
	"encoding/base64"
	"strings"

	"github.com/spf13/cast"
	duration "google.golang.org/protobuf/types/known/durationpb"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	wrappers "google.golang.org/protobuf/types/known/wrapperspb"
)

func ToTimestamp(v interface{}) (*timestamp.Timestamp, error) {
	if t, ok := v.(*timestamp.Timestamp); ok {
		return t, nil
	}
	if t, err := cast.ToTimeE(v); err != nil {
		return nil, err
	} else {
		return timestamp.New(t), nil
	}
}

func ToDuration(v interface{}) (*duration.Duration, error) {
	if d, ok := v.(*duration.Duration); ok {
		return d, nil
	}
	if d, err := cast.ToDurationE(v); err != nil {
		return nil, err
	} else {
		return duration.New(d), nil
	}
}

func ToDoubleWrapper(v interface{}) (*wrappers.DoubleValue, error) {
	if d, ok := v.(*wrappers.DoubleValue); ok {
		return d, nil
	}
	if d, err := cast.ToFloat64E(v); err != nil {
		return nil, err
	} else {
		return wrappers.Double(d), nil
	}
}

func ToFloatWrapper(v interface{}) (*wrappers.FloatValue, error) {
	if f, ok := v.(*wrappers.FloatValue); ok {
		return f, nil
	}
	if f, err := cast.ToFloat32E(v); err != nil {
		return nil, err
	} else {
		return wrappers.Float(f), nil
	}
}

func ToInt64Wrapper(v interface{}) (*wrappers.Int64Value, error) {
	if i, ok := v.(*wrappers.Int64Value); ok {
		return i, nil
	}
	if i, err := cast.ToInt64E(v); err != nil {
		return nil, err
	} else {
		return wrappers.Int64(i), nil
	}
}

func ToUInt64Wrapper(v interface{}) (*wrappers.UInt64Value, error) {
	if i, ok := v.(*wrappers.UInt64Value); ok {
		return i, nil
	}
	if i, err := cast.ToUint64E(v); err != nil {
		return nil, err
	} else {
		return wrappers.UInt64(i), nil
	}
}

func ToInt32Wrapper(v interface{}) (*wrappers.Int32Value, error) {
	if i, ok := v.(*wrappers.Int32Value); ok {
		return i, nil
	}
	if i, err := cast.ToInt32E(v); err != nil {
		return nil, err
	} else {
		return wrappers.Int32(i), nil
	}
}

func ToUInt32Wrapper(v interface{}) (*wrappers.UInt32Value, error) {
	if i, ok := v.(*wrappers.UInt32Value); ok {
		return i, nil
	}
	if i, err := cast.ToUint32E(v); err != nil {
		return nil, err
	} else {
		return wrappers.UInt32(i), nil
	}
}

func ToBoolWrapper(v interface{}) (*wrappers.BoolValue, error) {
	if b, ok := v.(*wrappers.BoolValue); ok {
		return b, nil
	}
	if b, err := cast.ToBoolE(v); err != nil {
		return nil, err
	} else {
		return wrappers.Bool(b), nil
	}
}

func ToStringWrapper(v interface{}) (*wrappers.StringValue, error) {
	if s, ok := v.(*wrappers.StringValue); ok {
		return s, nil
	}
	if s, err := cast.ToStringE(v); err != nil {
		return nil, err
	} else {
		return wrappers.String(s), nil
	}
}

func ToBytesWrapper(v interface{}) (*wrappers.BytesValue, error) {
	if b, ok := v.(*wrappers.BytesValue); ok {
		return b, nil
	}
	if b, ok := v.([]byte); ok {
		return wrappers.Bytes(b), nil
	}
	if s, err := cast.ToStringE(v); err != nil {
		return nil, err
	} else if v, err := base64.RawStdEncoding.DecodeString(strings.TrimRight(s, "=")); err != nil {
		return nil, err
	} else {
		return wrappers.Bytes(v), nil
	}
}
