package zerolog

import (
	"context"
	"io"
)

type Level int8

const (
	DebugLevel Level = iota

	InfoLevel

	WarnLevel

	ErrorLevel

	FatalLevel

	PanicLevel

	NoLevel

	Disabled

	TraceLevel Level = -1
)

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

func ParseLevel(levelStr string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

func (l *Level) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l Level) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Logger struct {
	w       LevelWriter
	level   Level
	sampler Sampler
	context []byte
	hooks   []Hook
	stack   bool
	ctx     context.Context
}

func New(w io.Writer) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func Nop() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l Logger) Output(w io.Writer) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l Logger) With() Context { _ = "STUB: not implemented"; return *new(Context) }

func (l *Logger) UpdateContext(update func(c Context) Context) { _ = "STUB: not implemented"; return }

func (l Logger) Level(lvl Level) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l Logger) GetLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

func (l Logger) Sample(s Sampler) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l Logger) Hook(hooks ...Hook) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l *Logger) Trace() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Debug() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Info() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Warn() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Error() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Err(err error) *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Fatal() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Panic() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) WithLevel(level Level) *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Log() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) Print(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Println(v ...interface{}) { _ = "STUB: not implemented"; return }

func (l Logger) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (l *Logger) newEvent(level Level, done func(string)) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (l *Logger) scratchEvent() *Event { _ = "STUB: not implemented"; return nil }

func (l *Logger) disabled() bool { _ = "STUB: not implemented"; return false }

func (l *Logger) should(lvl Level) bool { _ = "STUB: not implemented"; return false }
