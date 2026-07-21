//go:build !windows && !binary_log
// +build !windows,!binary_log

package zerolog

import (
	"io"
)

const ceePrefix = "@cee:"

type SyslogWriter interface {
	io.Writer
	Debug(m string) error
	Info(m string) error
	Warning(m string) error
	Err(m string) error
	Emerg(m string) error
	Crit(m string) error
}

type syslogWriter struct {
	w      SyslogWriter
	prefix string
}

func SyslogLevelWriter(w SyslogWriter) LevelWriter {
	_ = "STUB: not implemented"
	return *new(LevelWriter)
}

func SyslogCEEWriter(w SyslogWriter) LevelWriter {
	_ = "STUB: not implemented"
	return *new(LevelWriter)
}

func (sw syslogWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (sw syslogWriter) WriteLevel(level Level, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sw syslogWriter) Close() error { _ = "STUB: not implemented"; return nil }
