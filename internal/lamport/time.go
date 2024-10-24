package lamport

import (
	"sync/atomic"
)

type LamportClock struct {
	time atomic.Uint64
}

func (lc *LamportClock) Add() {
	lc.time.Add(1)
}
