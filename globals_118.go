//go:build go1.18
// +build go1.18

package zerolog

import (
	"fmt"
)

func AsLogObjectMarshalers[T LogObjectMarshaler](objs []T) []LogObjectMarshaler {
	_ = "STUB: not implemented"
	return nil
}

func AsStringers[T fmt.Stringer](objs []T) []fmt.Stringer { _ = "STUB: not implemented"; return nil }
