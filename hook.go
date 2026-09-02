package zerolog

type Hook interface {
	Run(e *Event, level Level, message string)
}

type HookFunc func(e *Event, level Level, message string)

func (h HookFunc) Run(e *Event, level Level, message string) { _ = "STUB: not implemented"; return }

type LevelHook struct {
	NoLevelHook, TraceHook, DebugHook, InfoHook, WarnHook, ErrorHook, FatalHook, PanicHook Hook
}

func (h LevelHook) Run(e *Event, level Level, message string) { _ = "STUB: not implemented"; return }

func NewLevelHook() LevelHook { _ = "STUB: not implemented"; return *new(LevelHook) }
