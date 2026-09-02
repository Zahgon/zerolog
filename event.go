package zerolog

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

var eventPool = &sync.Pool{
	New: func() interface{} {
		return &Event{
			buf: make([]byte, 0, 500),
		}
	},
}

type Event struct {
	buf       []byte
	w         LevelWriter
	level     Level
	done      func(msg string)
	stack     bool
	ch        []Hook
	skipFrame int
	ctx       context.Context
}

func putEvent(e *Event) { _ = "STUB: not implemented"; return }

type LogObjectMarshaler interface {
	MarshalZerologObject(e *Event)
}

type LogArrayMarshaler interface {
	MarshalZerologArray(a *Array)
}

func newEvent(w LevelWriter, level Level, stack bool, ctx context.Context, hooks []Hook) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) write() (err error) { _ = "STUB: not implemented"; return nil }

func (e *Event) Enabled() bool { _ = "STUB: not implemented"; return false }

func (e *Event) Discard() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Msg(msg string) { _ = "STUB: not implemented"; return }

func (e *Event) Send() { _ = "STUB: not implemented"; return }

func (e *Event) Msgf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (e *Event) MsgFunc(createMsg func() string) { _ = "STUB: not implemented"; return }

func (e *Event) msg(msg string) { _ = "STUB: not implemented"; return }

func (e *Event) Fields(fields interface{}) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Dict(key string, dict *Event) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) CreateDict() *Event { _ = "STUB: not implemented"; return nil }

func Dict() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) CreateArray() *Array { _ = "STUB: not implemented"; return nil }

func (e *Event) Array(key string, arr LogArrayMarshaler) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) appendObject(obj LogObjectMarshaler) { _ = "STUB: not implemented"; return }

func (e *Event) Object(key string, obj LogObjectMarshaler) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) Objects(key string, objs []LogObjectMarshaler) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) ObjectsV(key string, objs ...LogObjectMarshaler) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) Func(f func(e *Event)) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) EmbedObject(obj LogObjectMarshaler) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Str(key, val string) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Strs(key string, vals []string) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) StrsV(key string, vals ...string) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Stringer(key string, val fmt.Stringer) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) Stringers(key string, vals []fmt.Stringer) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) StringersV(key string, vals ...fmt.Stringer) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) Bytes(key string, val []byte) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Hex(key string, val []byte) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) RawJSON(key string, b []byte) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) RawCBOR(key string, b []byte) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) AnErr(key string, err error) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Errs(key string, errs []error) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Err(err error) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Stack() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ctx(ctx context.Context) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) GetCtx() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (e *Event) Bool(key string, b bool) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Bools(key string, b []bool) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Int(key string, i int) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ints(key string, i []int) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Int8(key string, i int8) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ints8(key string, i []int8) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Int16(key string, i int16) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ints16(key string, i []int16) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Int32(key string, i int32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ints32(key string, i []int32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Int64(key string, i int64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Ints64(key string, i []int64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uint(key string, i uint) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uints(key string, i []uint) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uint8(key string, i uint8) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uints8(key string, i []uint8) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uint16(key string, i uint16) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uints16(key string, i []uint16) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uint32(key string, i uint32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uints32(key string, i []uint32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uint64(key string, i uint64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Uints64(key string, i []uint64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Float32(key string, f float32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Floats32(key string, f []float32) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Float64(key string, f float64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Floats64(key string, f []float64) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Timestamp() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Time(key string, t time.Time) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Times(key string, t []time.Time) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Dur(key string, d time.Duration) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Durs(key string, d []time.Duration) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) TimeDiff(key string, t time.Time, start time.Time) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) Any(key string, i interface{}) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Interface(key string, i interface{}) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Type(key string, val interface{}) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) CallerSkipFrame(skip int) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Caller(skip ...int) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) caller(skip int) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) IPAddr(key string, ip net.IP) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) IPAddrs(key string, ip []net.IP) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) IPPrefix(key string, pfx net.IPNet) *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) IPPrefixes(key string, pfx []net.IPNet) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) MACAddr(key string, ha net.HardwareAddr) *Event {
	_ = "STUB: not implemented"
	return nil
}
