//go:build binary_log
// +build binary_log

package zerolog

import (
	"github.com/rs/zerolog/internal/cbor"
)

var (
	_ encoder = (*cbor.Encoder)(nil)

	enc = cbor.Encoder{}
)

func init() {

	cbor.JSONMarshalFunc = func(v interface{}) ([]byte, error) {
		return InterfaceMarshalFunc(v)
	}
}

func appendJSON(dst []byte, j []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendCBOR(dst []byte, c []byte) []byte { _ = "STUB: not implemented"; return nil }

func decodeIfBinaryToString(in []byte) string { _ = "STUB: not implemented"; return "" }

func decodeObjectToStr(in []byte) string { _ = "STUB: not implemented"; return "" }

func decodeIfBinaryToBytes(in []byte) []byte { _ = "STUB: not implemented"; return nil }
