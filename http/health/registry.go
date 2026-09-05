package health

import (
    "context"
    "sync"
    "time"
)

type Registry struct {
    mu      sync.RWMutex
    checks  map[string]Checker
    timeout time.Duration
}

func NewRegistry(opts ...Option) *Registry {
    r := &Registry{
        checks:  make(map[string]Checker),
        timeout: 5 * time.Second,
    }
    for _, opt := range opts {
        opt(r)
    }
    return r
}

type Option func(*Registry)

func WithTimeout(d time.Duration) Option {
    return func(r *Registry) {
        r.timeout = d
    }
}

func (r *Registry) Register(c Checker) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.checks[c.Name()] = c
}

func (r *Registry) RegisterFunc(name string, fn CheckFunc) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.checks[name] = fn
}

func (r *Registry) Run(ctx context.Context) map[string]*Result {
    r.mu.RLock()
    checks := make([]Checker, 0, len(r.checks))
    for _, c := range r.checks {
        checks = append(checks, c)
    }
    r.mu.RUnlock()

    results := make(map[string]*Result)
    var mu sync.Mutex
    var wg sync.WaitGroup

    for _, c := range checks {
        wg.Add(1)
        go func(check Checker) {
            defer wg.Done()
            ctx, cancel := context.WithTimeout(ctx, r.timeout)
            defer cancel()
            result, err := check.Check(ctx)
            if err != nil {
                result = &Result{
                    Status: StatusFail,
                    Message: err.Error(),
                    Time:   time.Now(),
                    Error:  err,
                }
            }
            mu.Lock()
            results[check.Name()] = result
            mu.Unlock()
        }(c)
    }
    wg.Wait()
    return results
}
