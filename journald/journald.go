//go:build !windows
// +build !windows

package journald

import (
	"io"

	"github.com/coreos/go-systemd/v22/journal"
)

const defaultJournalDPrio = journal.PriNotice

var SendFunc func(string, journal.Priority, map[string]string) error

func NewJournalDWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

type journalWriter struct {
}

func levelToJPrio(zLevel string) journal.Priority {
	_ = "STUB: not implemented"
	return *new(journal.Priority)
}

func sanitizeKey(key string) string { _ = "STUB: not implemented"; return "" }

func (w journalWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
