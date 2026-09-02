//go:build !binary_log
// +build !binary_log

package zerolog

import (
	"github.com/rs/zerolog/internal/json"
)

var (
	_ encoder = (*json.Encoder)(nil)

	enc = json.Encoder{}
)

func init() {

	json.JSONMarshalFunc = func(v interface{}) ([]byte, error) {
		return InterfaceMarshalFunc(v)
	}
}

func appendJSON(dst []byte, j []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendCBOR(dst []byte, cbor []byte) []byte { _ = "STUB: not implemented"; return nil }

func decodeIfBinaryToString(in []byte) string { _ = "STUB: not implemented"; return "" }

func decodeObjectToStr(in []byte) string { _ = "STUB: not implemented"; return "" }

func decodeIfBinaryToBytes(in []byte) []byte { _ = "STUB: not implemented"; return nil }
