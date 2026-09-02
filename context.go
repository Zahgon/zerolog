package zerolog

import (
	"context"
	"fmt"
	"math"
	"net"
	"time"
)

type Context struct {
	l Logger
}

func (c Context) Logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (c Context) Fields(fields interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Dict(key string, dict *Event) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) CreateDict() *Event { _ = "STUB: not implemented"; return nil }

func (c Context) CreateArray() *Array { _ = "STUB: not implemented"; return nil }

func (c Context) Array(key string, arr LogArrayMarshaler) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Object(key string, obj LogObjectMarshaler) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Objects(key string, objs []LogObjectMarshaler) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) ObjectsV(key string, objs ...LogObjectMarshaler) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) EmbedObject(obj LogObjectMarshaler) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Str(key, val string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Strs(key string, vals []string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) StrsV(key string, vals ...string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Stringer(key string, val fmt.Stringer) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Stringers(key string, vals []fmt.Stringer) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) StringersV(key string, vals ...fmt.Stringer) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Bytes(key string, val []byte) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Hex(key string, val []byte) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) RawJSON(key string, b []byte) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) AnErr(key string, err error) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Errs(key string, errs []error) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Err(err error) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Ctx(ctx context.Context) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Bool(key string, b bool) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Bools(key string, b []bool) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Int(key string, i int) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Ints(key string, i []int) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Int8(key string, i int8) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Ints8(key string, i []int8) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Int16(key string, i int16) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Ints16(key string, i []int16) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Int32(key string, i int32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Ints32(key string, i []int32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Int64(key string, i int64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Ints64(key string, i []int64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uint(key string, i uint) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Uints(key string, i []uint) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uint8(key string, i uint8) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uints8(key string, i []uint8) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uint16(key string, i uint16) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uints16(key string, i []uint16) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uint32(key string, i uint32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uints32(key string, i []uint32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uint64(key string, i uint64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Uints64(key string, i []uint64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Float32(key string, f float32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Floats32(key string, f []float32) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Float64(key string, f float64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Floats64(key string, f []float64) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

type timestampHook struct{}

func (ts timestampHook) Run(e *Event, level Level, msg string) { _ = "STUB: not implemented"; return }

var th = timestampHook{}

func (c Context) Timestamp() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) Time(key string, t time.Time) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Times(key string, t []time.Time) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Dur(key string, d time.Duration) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Durs(key string, d []time.Duration) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Interface(key string, i interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Type(key string, val interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Any(key string, i interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Reset() Context { _ = "STUB: not implemented"; return *new(Context) }

type callerHook struct {
	callerSkipFrameCount int
}

func newCallerHook(skipFrameCount int) callerHook {
	_ = "STUB: not implemented"
	return *new(callerHook)
}

func (ch callerHook) Run(e *Event, level Level, msg string) { _ = "STUB: not implemented"; return }

const useGlobalSkipFrameCount = math.MinInt32

var ch = newCallerHook(useGlobalSkipFrameCount)

func (c Context) Caller() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) CallerWithSkipFrameCount(skipFrameCount int) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) Stack() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c Context) IPAddr(key string, ip net.IP) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) IPAddrs(key string, ip []net.IP) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) IPPrefix(key string, pfx net.IPNet) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) IPPrefixes(key string, pfx []net.IPNet) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c Context) MACAddr(key string, ha net.HardwareAddr) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}
