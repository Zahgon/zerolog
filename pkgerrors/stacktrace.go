package pkgerrors

import (
	"github.com/pkg/errors"
)

var (
	StackSourceFileName     = "source"
	StackSourceLineName     = "line"
	StackSourceFunctionName = "func"
)

type state struct {
	b []byte
}

func (s *state) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *state) Width() (wid int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func (s *state) Precision() (prec int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func (s *state) Flag(c int) bool { _ = "STUB: not implemented"; return false }

func frameField(f errors.Frame, s *state, c rune) string { _ = "STUB: not implemented"; return "" }

func MarshalStack(err error) interface{} { _ = "STUB: not implemented"; return nil }
