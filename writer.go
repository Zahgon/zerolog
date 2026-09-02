package zerolog

import (
	"bytes"
	"io"
	"sync"
)

type LevelWriter interface {
	io.Writer
	WriteLevel(level Level, p []byte) (n int, err error)
}

type LevelWriterAdapter struct {
	io.Writer
}

func (lw LevelWriterAdapter) WriteLevel(l Level, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (lw LevelWriterAdapter) Close() error { _ = "STUB: not implemented"; return nil }

type syncWriter struct {
	mu sync.Mutex
	lw LevelWriter
}

func SyncWriter(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (s *syncWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *syncWriter) WriteLevel(l Level, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *syncWriter) Close() error { _ = "STUB: not implemented"; return nil }

type multiLevelWriter struct {
	writers []LevelWriter
}

func (t multiLevelWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t multiLevelWriter) WriteLevel(l Level, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t multiLevelWriter) Close() error { _ = "STUB: not implemented"; return nil }

func MultiLevelWriter(writers ...io.Writer) LevelWriter {
	_ = "STUB: not implemented"
	return *new(LevelWriter)
}

type TestingLog interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Helper()
}

type TestWriter struct {
	T TestingLog

	Frame int
}

func NewTestWriter(t TestingLog) TestWriter { _ = "STUB: not implemented"; return *new(TestWriter) }

func (t TestWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func ConsoleTestWriter(t TestingLog) func(w *ConsoleWriter) { _ = "STUB: not implemented"; return nil }

type FilteredLevelWriter struct {
	Writer LevelWriter
	Level  Level
}

func (w *FilteredLevelWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *FilteredLevelWriter) WriteLevel(level Level, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *FilteredLevelWriter) Close() error { _ = "STUB: not implemented"; return nil }

var triggerWriterPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

type TriggerLevelWriter struct {
	io.Writer

	ConditionalLevel Level

	TriggerLevel Level

	buf       *bytes.Buffer
	triggered bool
	mu        sync.Mutex
}

func (w *TriggerLevelWriter) WriteLevel(l Level, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *TriggerLevelWriter) trigger() error { _ = "STUB: not implemented"; return nil }

func (w *TriggerLevelWriter) Trigger() error { _ = "STUB: not implemented"; return nil }

func (w *TriggerLevelWriter) Close() error { _ = "STUB: not implemented"; return nil }
