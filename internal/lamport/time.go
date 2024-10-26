package lamport

import (
	"sync/atomic"
)

type Clock struct {
	time atomic.Uint64
}

func (c *Clock) Now() uint64 {
	return c.time.Load()
}

func (c *Clock) Step() {
	c.time.Add(1)
}

func (c *Clock) Max(other uint64) {
	if c.time.Load() < other {
		c.time.Swap(other)
	}
}
