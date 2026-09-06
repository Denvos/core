package ratelimit

import (
    "net/http"
    "sync"
)

type Middleware struct {
    limiter Limiter
    keyFunc func(*http.Request) string
    store   map[string]Limiter
    mu      sync.RWMutex
    config  Config
}

func NewMiddleware(cfg Config) *Middleware {
    m := &Middleware{
        store:  make(map[string]Limiter),
        config: cfg,
        keyFunc: func(r *http.Request) string {
            return r.RemoteAddr
        },
    }
    return m
}

func (m *Middleware) KeyFunc(fn func(*http.Request) string) *Middleware {
    m.keyFunc = fn
    return m
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        key := m.keyFunc(r)
        m.mu.Lock()
        limiter, ok := m.store[key]
        if !ok {
            limiter = NewTokenBucket(m.config)
            m.store[key] = limiter
        }
        m.mu.Unlock()
        if !limiter.Allow() {
            w.Header().Set("Retry-After", "1")
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func (m *Middleware) Clear() {
    m.mu.Lock()
    defer m.mu.Unlock()
    for _, l := range m.store {
        if closer, ok := l.(interface{ Close() }); ok {
            closer.Close()
        }
    }
    m.store = make(map[string]Limiter)
}
