package memcached

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	client *memcache.Client
}

func New(addr string) *Cache {
	return &Cache{client: memcache.New(addr)}
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, bool, error) {
	item, err := c.client.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return item.Value, true, nil
}

func (c *Cache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Seconds()),
	})
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Delete(key)
}

func (c *Cache) Clear(ctx context.Context) error {
	return c.client.FlushAll()
}
