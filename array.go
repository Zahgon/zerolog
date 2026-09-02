package zerolog

import (
	"context"
	"net"
	"sync"
	"time"
)

var arrayPool = &sync.Pool{
	New: func() interface{} {
		return &Array{
			buf: make([]byte, 0, 500),
		}
	},
}

type Array struct {
	buf   []byte
	stack bool
	ctx   context.Context
	ch    []Hook
}

func putArray(a *Array) { _ = "STUB: not implemented"; return }

func Arr() *Array { _ = "STUB: not implemented"; return nil }

func (*Array) MarshalZerologArray(*Array) { _ = "STUB: not implemented"; return }

func (a *Array) write(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (a *Array) Object(obj LogObjectMarshaler) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Str(val string) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Bytes(val []byte) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Hex(val []byte) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) RawJSON(val []byte) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Err(err error) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Errs(errs []error) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Bool(b bool) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Int(i int) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Int8(i int8) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Int16(i int16) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Int32(i int32) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Int64(i int64) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Uint(i uint) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Uint8(i uint8) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Uint16(i uint16) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Uint32(i uint32) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Uint64(i uint64) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Float32(f float32) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Float64(f float64) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Time(t time.Time) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Dur(d time.Duration) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Interface(i interface{}) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) IPAddr(ip net.IP) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) IPPrefix(pfx net.IPNet) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) MACAddr(ha net.HardwareAddr) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Dict(dict *Event) *Array { _ = "STUB: not implemented"; return nil }

func (a *Array) Type(val interface{}) *Array { _ = "STUB: not implemented"; return nil }
