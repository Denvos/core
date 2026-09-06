package ratelimit

import (
    "sync"
    "time"
)

type Limiter interface {
    Allow() bool
    Wait()
    Reset()
}

type Config struct {
    Rate  int
    Burst int
    TTL   time.Duration
}

type TokenBucket struct {
    mu       sync.Mutex
    rate     int
    burst    int
    tokens   int
    last     time.Time
    ttl      time.Duration
    stop     chan struct{}
    once     sync.Once
}

func NewTokenBucket(cfg Config) *TokenBucket {
    if cfg.Rate <= 0 {
        cfg.Rate = 10
    }
    if cfg.Burst <= 0 {
        cfg.Burst = cfg.Rate
    }
    if cfg.TTL <= 0 {
        cfg.TTL = time.Minute
    }
    tb := &TokenBucket{
        rate:   cfg.Rate,
        burst:  cfg.Burst,
        tokens: cfg.Burst,
        last:   time.Now(),
        ttl:    cfg.TTL,
        stop:   make(chan struct{}),
    }
    go tb.refill()
    return tb
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    if tb.tokens > 0 {
        tb.tokens--
        return true
    }
    return false
}

func (tb *TokenBucket) Wait() {
    for !tb.Allow() {
        time.Sleep(10 * time.Millisecond)
    }
}

func (tb *TokenBucket) Reset() {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    tb.tokens = tb.burst
    tb.last = time.Now()
}

func (tb *TokenBucket) refill() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            tb.mu.Lock()
            now := time.Now()
            elapsed := now.Sub(tb.last)
            tb.last = now
            add := int(elapsed.Seconds() * float64(tb.rate))
            if add > 0 {
                tb.tokens += add
                if tb.tokens > tb.burst {
                    tb.tokens = tb.burst
                }
            }
            tb.mu.Unlock()
        case <-tb.stop:
            return
        }
    }
}

func (tb *TokenBucket) Close() {
    tb.once.Do(func() {
        close(tb.stop)
    })
}
