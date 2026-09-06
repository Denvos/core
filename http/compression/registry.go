package compression

import (
    "sync"
)

type Registry struct {
    mu      sync.RWMutex
    codecs  map[string]Compressor
    defaultOrder []string
}

func NewRegistry() *Registry {
    r := &Registry{
        codecs:       make(map[string]Compressor),
        defaultOrder: []string{},
    }
    // Register default compressors
    r.Register(NewGzip())
    r.Register(NewZstd())
    r.Register(NewBrotli())
    return r
}

func (r *Registry) Register(c Compressor) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.codecs[c.Name()] = c
    // Append to order if not already present
    for _, name := range r.defaultOrder {
        if name == c.Name() {
            return
        }
    }
    r.defaultOrder = append(r.defaultOrder, c.Name())
}

func (r *Registry) Get(name string) Compressor {
    r.mu.RLock()
    defer r.mu.RUnlock()
    return r.codecs[name]
}

func (r *Registry) List() []string {
    r.mu.RLock()
    defer r.mu.RUnlock()
    names := make([]string, 0, len(r.codecs))
    for name := range r.codecs {
        names = append(names, name)
    }
    return names
}

func (r *Registry) Negotiate(acceptEncoding string) Compressor {
    r.mu.RLock()
    defer r.mu.RUnlock()

    // Parse Accept-Encoding header (simplified)
    // Example: "gzip, deflate, br;q=0.9"
    if acceptEncoding == "" {
        return nil
    }
    // For now, just check each in order
    for _, name := range r.defaultOrder {
        c := r.codecs[name]
        if c != nil && c.Match(acceptEncoding) {
            return c
        }
    }
    return nil
}
