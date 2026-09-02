package zerolog

import (
	"time"
)

var (
	Often = RandomSampler(10)

	Sometimes = RandomSampler(100)

	Rarely = RandomSampler(1000)
)

type Sampler interface {
	Sample(lvl Level) bool
}

type RandomSampler uint32

func (s RandomSampler) Sample(lvl Level) bool { _ = "STUB: not implemented"; return false }

type BasicSampler struct {
	N       uint32
	counter uint32
}

func (s *BasicSampler) Sample(lvl Level) bool { _ = "STUB: not implemented"; return false }

type BurstSampler struct {
	Burst uint32

	Period time.Duration

	NextSampler Sampler

	counter uint32
	resetAt int64
}

func (s *BurstSampler) Sample(lvl Level) bool { _ = "STUB: not implemented"; return false }

func (s *BurstSampler) inc() uint32 { _ = "STUB: not implemented"; return 0 }

type LevelSampler struct {
	TraceSampler, DebugSampler, InfoSampler, WarnSampler, ErrorSampler Sampler
}

func (s LevelSampler) Sample(lvl Level) bool { _ = "STUB: not implemented"; return false }
