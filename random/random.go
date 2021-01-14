package random

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

// It returns random value in [a..b).
func IntForInterval(r *rand.Rand, a, b int) int {
	switch {
	case a < b:
		return a + r.Intn(b-a)
	case a > b:
		return b + r.Intn(a-b)
	default: // a = b
		return a
	}
}
