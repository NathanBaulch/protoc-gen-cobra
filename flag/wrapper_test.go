package flag

import (
	"bytes"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	//"github.com/golang/protobuf/ptypes/wrappers"
	wrappers "google.golang.org/protobuf/types/known/wrapperspb"
)

func TestBoolWrapperVar(t *testing.T) {
	testCases := []struct {
		val  string
		want bool
	}{
		{"0", false},
		{"1", true},
		{"false", false},
		{"true", true},
	}
	for _, tc := range testCases {
		fs := &pflag.FlagSet{}
		var v *wrappers.BoolValue
		BoolWrapperVar(fs, &v, "foo", "")
		assert.NoError(t, fs.Set("foo", tc.val))
		assert.NotNil(t, v)
		assert.Equal(t, v.Value, tc.want)
	}
}

func TestBytesBase64WrapperVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v *wrappers.BytesValue
	BytesBase64WrapperVar(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.NotNil(t, v)
	assert.Equal(t, "abcd", string(v.Value))
}

func TestBytesBase64WrapperVarStdIn(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v *wrappers.BytesValue
	BytesBase64WrapperVar(fs, &v, "foo", "")
	stdin = bytes.NewReader([]byte("abcd"))
	assert.NoError(t, fs.Set("foo", "-"))
	assert.NotNil(t, v)
	assert.Equal(t, "abcd", string(v.Value))
}

func TestBytesBase64WrapperSliceVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v []*wrappers.BytesValue
	BytesBase64WrapperSliceVar(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.NoError(t, fs.Set("foo", "MTIzNA=="))
	assert.Len(t, v, 2)
	assert.Equal(t, "abcd", string(v[0].Value))
	assert.Equal(t, "1234", string(v[1].Value))
}
