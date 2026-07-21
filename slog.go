package zerolog

import (
	"context"
	"log/slog"
)

type SlogHandler struct {
	logger Logger
	prefix string
	attrs  []slog.Attr
}

func NewSlogHandler(logger Logger) *SlogHandler { _ = "STUB: not implemented"; return nil }

func (h *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *SlogHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *SlogHandler) hasTimestampHook() bool { _ = "STUB: not implemented"; return false }

func (h *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *SlogHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *SlogHandler) clone() *SlogHandler { _ = "STUB: not implemented"; return nil }

func slogToZerologLevel(level slog.Level) Level { _ = "STUB: not implemented"; return *new(Level) }

func zerologToSlogLevel(level Level) slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }

func joinPrefix(prefix, key string) string { _ = "STUB: not implemented"; return "" }

func appendSlogAttr(event *Event, attr slog.Attr, prefix string) *Event {
	_ = "STUB: not implemented"
	return nil
}

var _ slog.Handler = (*SlogHandler)(nil)
