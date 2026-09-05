package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
	Stats() Stats
	Close() error
}

type Stats struct {
	Hits   int64
	Misses int64
	Size   int64
	Evictions int64
}

type Option func(interface{})

type Config struct {
	TTL      time.Duration
	MaxSize  int64
	Eviction string // lru, lfu, fifo
}

var DefaultConfig = Config{
	TTL:      5 * time.Minute,
	MaxSize:  1000,
	Eviction: "lru",
}
