package ratelimit

import (
    "sync"
    "time"
)

type SlidingWindow struct {
    mu      sync.Mutex
    rate    int
    window  time.Duration
    entries []time.Time
}

func NewSlidingWindow(rate int, window time.Duration) *SlidingWindow {
    if rate <= 0 {
        rate = 10
    }
    if window <= 0 {
        window = time.Second
    }
    return &SlidingWindow{
        rate:    rate,
        window:  window,
        entries: make([]time.Time, 0, rate),
    }
}

func (sw *SlidingWindow) Allow() bool {
    sw.mu.Lock()
    defer sw.mu.Unlock()
    now := time.Now()
    cutoff := now.Add(-sw.window)
    i := 0
    for i < len(sw.entries) && sw.entries[i].Before(cutoff) {
        i++
    }
    if i > 0 {
        sw.entries = sw.entries[i:]
    }
    if len(sw.entries) >= sw.rate {
        return false
    }
    sw.entries = append(sw.entries, now)
    return true
}

func (sw *SlidingWindow) Wait() {
    for !sw.Allow() {
        time.Sleep(10 * time.Millisecond)
    }
}

func (sw *SlidingWindow) Reset() {
    sw.mu.Lock()
    defer sw.mu.Unlock()
    sw.entries = sw.entries[:0]
}
