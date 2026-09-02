package zerolog

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

const (
	TimeFormatUnix = ""

	TimeFormatUnixMs = "UNIXMS"

	TimeFormatUnixMicro = "UNIXMICRO"

	TimeFormatUnixNano = "UNIXNANO"

	DurationFormatFloat = "float"

	DurationFormatInt = "int"

	DurationFormatString = "string"
)

var (
	TimestampFieldName = "time"

	LevelFieldName = "level"

	LevelTraceValue = "trace"

	LevelDebugValue = "debug"

	LevelInfoValue = "info"

	LevelWarnValue = "warn"

	LevelErrorValue = "error"

	LevelFatalValue = "fatal"

	LevelPanicValue = "panic"

	LevelFieldMarshalFunc = func(l Level) string {
		return l.String()
	}

	MessageFieldName = "message"

	ErrorFieldName = "error"

	CallerFieldName = "caller"

	CallerSkipFrameCount = 2

	CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return file + ":" + strconv.Itoa(line)
	}

	ErrorStackFieldName = "stack"

	ErrorStackMarshaler func(err error) interface{}

	ErrorMarshalFunc = func(err error) interface{} {
		return err
	}

	InterfaceMarshalFunc = func(v interface{}) ([]byte, error) {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(v)
		if err != nil {
			return nil, err
		}
		b := buf.Bytes()
		if len(b) > 0 {

			return b[:len(b)-1], nil
		}
		return b, nil
	}

	TimeFieldFormat = time.RFC3339

	TimestampFunc = time.Now

	DurationFieldFormat = DurationFormatFloat

	DurationFieldUnit = time.Millisecond

	DurationFieldInteger = false

	ErrorHandler func(err error)

	FatalExitFunc func()

	DefaultContextLogger *Logger

	LevelColors = map[Level]int{
		TraceLevel: colorBlue,
		DebugLevel: 0,
		InfoLevel:  colorGreen,
		WarnLevel:  colorYellow,
		ErrorLevel: colorRed,
		FatalLevel: colorRed,
		PanicLevel: colorRed,
	}

	FormattedLevels = map[Level]string{
		TraceLevel: "TRC",
		DebugLevel: "DBG",
		InfoLevel:  "INF",
		WarnLevel:  "WRN",
		ErrorLevel: "ERR",
		FatalLevel: "FTL",
		PanicLevel: "PNC",
	}

	TriggerLevelWriterBufferReuseLimit = 64 * 1024

	FloatingPointPrecision = -1
)

var (
	gLevel          = new(int32)
	disableSampling = new(int32)
)

func SetGlobalLevel(l Level) { _ = "STUB: not implemented"; return }

func GlobalLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

func DisableSampling(v bool) { _ = "STUB: not implemented"; return }

func samplingDisabled() bool { _ = "STUB: not implemented"; return false }
