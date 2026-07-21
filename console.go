package zerolog

import (
	"bytes"
	"io"
	"sync"
	"time"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90

	unknownLevel = "???"
)

var (
	consoleBufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 100))
		},
	}
)

const (
	consoleDefaultTimeFormat = time.Kitchen
)

type Formatter func(interface{}) string

type FormatterByFieldName func(interface{}, string) string

type ConsoleWriter struct {
	Out io.Writer

	NoColor bool

	TimeFormat string

	TimeLocation *time.Location

	PartsOrder []string

	PartsExclude []string

	FieldsOrder []string

	fieldIsOrdered map[string]int

	FieldsExclude []string

	FormatTimestamp     Formatter
	FormatLevel         Formatter
	FormatCaller        Formatter
	FormatMessage       Formatter
	FormatFieldName     Formatter
	FormatFieldValue    Formatter
	FormatErrFieldName  Formatter
	FormatErrFieldValue Formatter

	FormatPartValueByName FormatterByFieldName

	FormatExtra func(map[string]interface{}, *bytes.Buffer) error

	FormatPrepare func(map[string]interface{}) error
}

func NewConsoleWriter(options ...func(w *ConsoleWriter)) ConsoleWriter {
	_ = "STUB: not implemented"
	return *new(ConsoleWriter)
}

func (w ConsoleWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w ConsoleWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w ConsoleWriter) writeFields(evt map[string]interface{}, buf *bytes.Buffer) {
	_ = "STUB: not implemented"
	return
}

func (w ConsoleWriter) writePart(buf *bytes.Buffer, evt map[string]interface{}, p string) {
	_ = "STUB: not implemented"
	return
}

func (w ConsoleWriter) orderFields(fields []string) { _ = "STUB: not implemented"; return }

func needsQuote(s string) bool { _ = "STUB: not implemented"; return false }

func colorize(s interface{}, c int, disabled bool) string { _ = "STUB: not implemented"; return "" }

func consoleDefaultPartsOrder() []string { _ = "STUB: not implemented"; return nil }

func consoleDefaultFormatTimestamp(timeFormat string, location *time.Location, noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func stripLevel(ll string) string { _ = "STUB: not implemented"; return "" }

func consoleDefaultFormatLevel(noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func consoleDefaultFormatCaller(noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func consoleDefaultFormatMessage(noColor bool, level interface{}) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func consoleDefaultFormatFieldName(noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func consoleDefaultFormatFieldValue(i interface{}) string { _ = "STUB: not implemented"; return "" }

func consoleDefaultFormatErrFieldName(noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func consoleDefaultFormatErrFieldValue(noColor bool) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}
