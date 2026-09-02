package diode

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/rs/zerolog/diode/internal/diodes"
)

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 500)
	},
}

type Alerter func(missed int)

type diodeFetcher interface {
	diodes.Diode
	Next() diodes.GenericDataType
}

type Writer struct {
	w    io.Writer
	d    diodeFetcher
	c    context.CancelFunc
	done chan struct{}
}

func NewWriter(w io.Writer, size int, pollInterval time.Duration, f Alerter) Writer {
	_ = "STUB: not implemented"
	return *new(Writer)
}

func (dw Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (dw Writer) Close() error { _ = "STUB: not implemented"; return nil }

func (dw Writer) poll() { _ = "STUB: not implemented"; return }
