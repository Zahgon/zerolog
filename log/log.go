package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func Output(w io.Writer) zerolog.Logger { _ = "STUB: not implemented"; return *new(zerolog.Logger) }

func With() zerolog.Context { _ = "STUB: not implemented"; return *new(zerolog.Context) }

func Level(level zerolog.Level) zerolog.Logger {
	_ = "STUB: not implemented"
	return *new(zerolog.Logger)
}

func Sample(s zerolog.Sampler) zerolog.Logger {
	_ = "STUB: not implemented"
	return *new(zerolog.Logger)
}

func Hook(h zerolog.Hook) zerolog.Logger { _ = "STUB: not implemented"; return *new(zerolog.Logger) }

func Err(err error) *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Trace() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Debug() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Info() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Warn() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Error() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Fatal() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Panic() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func WithLevel(level zerolog.Level) *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Log() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func Print(v ...interface{}) { _ = "STUB: not implemented"; return }

func Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Ctx(ctx context.Context) *zerolog.Logger { _ = "STUB: not implemented"; return nil }
