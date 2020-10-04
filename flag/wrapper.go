package flag

import (
	"github.com/spf13/pflag"
	wrappers "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/gutterbacon/protoc-gen-cobra/ptypes"
)

func BoolWrapperVar(fs *pflag.FlagSet, p **wrappers.BoolValue, name, usage string) {
	v := fs.String(name, "", usage)
	WithPostSetHookE(fs, name, func() (err error) { *p, err = ptypes.ToBoolWrapper(v); return })
}

func BoolWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.BoolValue, name, usage string) {
	v := fs.BoolSlice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.BoolValue, len(*v))
		for i, item := range *v {
			out[i] = wrappers.Bool(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseBoolWrapper(val string) (interface{}, error) { return ptypes.ToBoolWrapper(val) }

func Int32WrapperVar(fs *pflag.FlagSet, p **wrappers.Int32Value, name, usage string) {
	v := fs.Int32(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.Int32(*v) })
}

func Int32WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.Int32Value, name, usage string) {
	v := fs.Int32Slice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.Int32Value, len(*v))
		for i, item := range *v {
			out[i] = wrappers.Int32(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseInt32Wrapper(val string) (interface{}, error) { return ptypes.ToInt32Wrapper(val) }

func Int64WrapperVar(fs *pflag.FlagSet, p **wrappers.Int64Value, name, usage string) {
	v := fs.Int64(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.Int64(*v) })
}

func Int64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.Int64Value, name, usage string) {
	v := fs.Int64Slice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.Int64Value, len(*v))
		for i, item := range *v {
			out[i] = wrappers.Int64(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseInt64Wrapper(val string) (interface{}, error) { return ptypes.ToInt64Wrapper(val) }

func UInt32WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt32Value, name, usage string) {
	v := fs.Uint32(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.UInt32(*v) })
}

func UInt32WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.UInt32Value, name, usage string) {
	var v []uint32
	Uint32SliceVar(fs, &v, name, usage)
	hook := func() {
		out := make([]*wrappers.UInt32Value, len(v))
		for i, item := range v {
			out[i] = wrappers.UInt32(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseUInt32Wrapper(val string) (interface{}, error) { return ptypes.ToUInt32Wrapper(val) }

func UInt64WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt64Value, name, usage string) {
	v := fs.Uint64(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.UInt64(*v) })
}

func UInt64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.UInt64Value, name, usage string) {
	var v []uint64
	Uint64SliceVar(fs, &v, name, usage)
	hook := func() {
		out := make([]*wrappers.UInt64Value, len(v))
		for i, item := range v {
			out[i] = wrappers.UInt64(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseUInt64Wrapper(val string) (interface{}, error) { return ptypes.ToUInt64Wrapper(val) }

func FloatWrapperVar(fs *pflag.FlagSet, p **wrappers.FloatValue, name, usage string) {
	v := fs.Float32(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.Float(*v) })
}

func FloatWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.FloatValue, name, usage string) {
	v := fs.Float32Slice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.FloatValue, len(*v))
		for i, item := range *v {
			out[i] = wrappers.Float(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseFloatWrapper(val string) (interface{}, error) { return ptypes.ToFloatWrapper(val) }

func DoubleWrapperVar(fs *pflag.FlagSet, p **wrappers.DoubleValue, name, usage string) {
	v := fs.Float64(name, 0, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.Double(*v) })
}

func DoubleWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.DoubleValue, name, usage string) {
	v := fs.Float64Slice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.DoubleValue, len(*v))
		for i, item := range *v {
			out[i] = wrappers.Double(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseDoubleWrapper(val string) (interface{}, error) { return ptypes.ToDoubleWrapper(val) }

func StringWrapperVar(fs *pflag.FlagSet, p **wrappers.StringValue, name, usage string) {
	v := fs.String(name, "", usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.String(*v) })
}

func StringWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.StringValue, name, usage string) {
	v := fs.StringSlice(name, nil, usage)
	hook := func() {
		out := make([]*wrappers.StringValue, len(*v))
		for i, item := range *v {
			out[i] = wrappers.String(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseStringWrapper(val string) (interface{}, error) { return ptypes.ToStringWrapper(val) }

func BytesBase64WrapperVar(fs *pflag.FlagSet, p **wrappers.BytesValue, name, usage string) {
	var v []byte
	BytesBase64Var(fs, &v, name, usage)
	WithPostSetHook(fs, name, func() { *p = wrappers.Bytes(v) })
}

func BytesBase64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.BytesValue, name, usage string) {
	var v [][]byte
	BytesBase64SliceVar(fs, &v, name, usage)
	hook := func() {
		out := make([]*wrappers.BytesValue, len(v))
		for i, item := range v {
			out[i] = wrappers.Bytes(item)
		}
		*p = out
	}
	WithPostSetHook(fs, name, hook)
}

func ParseBytesBase64Wrapper(val string) (interface{}, error) { return ptypes.ToBytesWrapper(val) }

func WithPostSetHook(fs *pflag.FlagSet, name string, hook func()) {
	WithPostSetHookE(fs, name, func() error { hook(); return nil })
}

func WithPostSetHookE(fs *pflag.FlagSet, name string, hook func() error) {
	f := fs.Lookup(name)
	f.Value = &postSetHookValue{f.Value, hook}
}

type postSetHookValue struct {
	pflag.Value
	hook func() error
}

func (v *postSetHookValue) Set(s string) error {
	if err := v.Value.Set(s); err != nil {
		return err
	}
	return v.hook()
}
