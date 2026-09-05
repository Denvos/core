package cache

import (
	"testing"
	"time"

	"github.com/Denvos/core/cache/ttl"
	"github.com/Denvos/core/cache/lru"
)

func TestTTL(t *testing.T) {
	c := ttl.New(100 * time.Millisecond)
	c.Set("key", "value", 0)
	if val, ok := c.Get("key"); !ok || val != "value" {
		t.Error("expected value")
	}
	time.Sleep(150 * time.Millisecond)
	if _, ok := c.Get("key"); ok {
		t.Error("expected key to expire")
	}
}

func TestLRU(t *testing.T) {
	c := lru.New(lru.WithMaxSize(2))
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error("expected a to be evicted")
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("expected b to exist")
	}
}
