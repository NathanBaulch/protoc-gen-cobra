package iocodec

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestDecodeValue(t *testing.T) {
	v := map[string]interface{}{
		"timestamp": time.Unix(123, 456).UTC(),
		"duration":  time.Duration(123),
		"bool":      true,
		"bytes":     []byte("abc"),
		"double":    123.456,
		"float":     float32(123.456),
		"int32":     int32(123),
		"uint32":    uint32(123),
		"int64":     int64(123),
		"uint64":    uint64(123),
		"string":    "abc",
	}
	p := &struct {
		Timestamp *timestamppb.Timestamp
		Duration  *durationpb.Duration
		Bool      *wrapperspb.BoolValue
		Bytes     *wrapperspb.BytesValue
		Double    *wrapperspb.DoubleValue
		Float     *wrapperspb.FloatValue
		Int32     *wrapperspb.Int32Value
		UInt32    *wrapperspb.UInt32Value
		Int64     *wrapperspb.Int64Value
		UInt64    *wrapperspb.UInt64Value
		String    *wrapperspb.StringValue
	}{}
	assert.NoError(t, decodeValue(v, p))
	assert.Equal(t, v["timestamp"], p.Timestamp.AsTime())
	assert.Equal(t, v["duration"], p.Duration.AsDuration())
	assert.Equal(t, v["bool"], p.Bool.Value)
	assert.Equal(t, v["bytes"], p.Bytes.Value)
	assert.Equal(t, v["double"], p.Double.Value)
	assert.Equal(t, v["float"], p.Float.Value)
	assert.Equal(t, v["int32"], p.Int32.Value)
	assert.Equal(t, v["uint32"], p.UInt32.Value)
	assert.Equal(t, v["int64"], p.Int64.Value)
	assert.Equal(t, v["uint64"], p.UInt64.Value)
	assert.Equal(t, v["string"], p.String.Value)
}

func TestEncodeValue(t *testing.T) {
	type Foo struct {
		Bar interface{} `json:"bar"`
	}
	testCases := []struct {
		val  interface{}
		want interface{}
	}{
		{timestamppb.New(time.Unix(123, 456)), time.Unix(123, 456).UTC()},
		{durationpb.New(time.Duration(123)), time.Duration(123)},
		{wrapperspb.Bool(true), true},
		{wrapperspb.Bytes([]byte("abc")), []byte("abc")},
		{wrapperspb.Double(123.456), 123.456},
		{wrapperspb.Float(123.456), float32(123.456)},
		{wrapperspb.Int32(123), int32(123)},
		{wrapperspb.UInt32(123), uint32(123)},
		{wrapperspb.Int64(123), int64(123)},
		{wrapperspb.UInt64(123), uint64(123)},
		{wrapperspb.String("abc"), "abc"},
		{(*int)(nil), (*int)(nil)},
	}
	for _, tc := range testCases {
		if got, err := encodeValue(tc.val); true {
			if tc.val == (*int)(nil) {
				assert.Error(t, errUnchanged)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, tc.want, got)
			}
		}

		if got, err := encodeValue(Foo{Bar: tc.val}); true {
			assert.NoError(t, err)
			assert.EqualValues(t, map[string]interface{}{"bar": tc.want}, got)
		}

		if got, err := encodeValue(map[string]interface{}{"foo": tc.val}); true {
			assert.NoError(t, err)
			assert.EqualValues(t, map[string]interface{}{"foo": tc.want}, got)
		}

		if got, err := encodeValue([]interface{}{tc.val}); true {
			assert.NoError(t, err)
			assert.EqualValues(t, []interface{}{tc.want}, got)
		}
	}
}
