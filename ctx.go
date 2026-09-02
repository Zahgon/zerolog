package zerolog

import (
	"context"
)

var disabledLogger *Logger

func init() {
	SetGlobalLevel(TraceLevel)
	l := Nop()
	disabledLogger = &l
}

type ctxKey struct{}

func (l Logger) WithContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Ctx(ctx context.Context) *Logger { _ = "STUB: not implemented"; return nil }
