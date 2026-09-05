package stats

import "sync/atomic"

type Stats struct {
	Hits      int64
	Misses    int64
	Size      int64
	Evictions int64
	Sets      int64
	Deletes   int64
}

func (s *Stats) Hit() {
	atomic.AddInt64(&s.Hits, 1)
}

func (s *Stats) Miss() {
	atomic.AddInt64(&s.Misses, 1)
}

func (s *Stats) Set() {
	atomic.AddInt64(&s.Sets, 1)
}

func (s *Stats) Delete() {
	atomic.AddInt64(&s.Deletes, 1)
}

func (s *Stats) Evict() {
	atomic.AddInt64(&s.Evictions, 1)
}

func (s *Stats) AddSize(delta int64) {
	atomic.AddInt64(&s.Size, delta)
}

func (s *Stats) Snapshot() map[string]int64 {
	return map[string]int64{
		"hits":      atomic.LoadInt64(&s.Hits),
		"misses":    atomic.LoadInt64(&s.Misses),
		"size":      atomic.LoadInt64(&s.Size),
		"evictions": atomic.LoadInt64(&s.Evictions),
		"sets":      atomic.LoadInt64(&s.Sets),
		"deletes":   atomic.LoadInt64(&s.Deletes),
	}
}
