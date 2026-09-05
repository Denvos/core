package lru

import (
	"container/list"
	"sync"
)

type Cache struct {
	mu       sync.RWMutex
	items    map[string]*list.Element
	order    *list.List
	maxSize  int64
	size     int64
	evictions int64
}

type entry struct {
	key   string
	value interface{}
}

type Option func(*Cache)

func WithMaxSize(max int64) Option {
	return func(c *Cache) {
		c.maxSize = max
	}
}

func New(opts ...Option) *Cache {
	c := &Cache{
		items:   make(map[string]*list.Element),
		order:   list.New(),
		maxSize: 1000,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.order.MoveToFront(el)
	return el.Value.(*entry).value, true
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		el.Value.(*entry).value = value
		return
	}
	if c.maxSize > 0 && c.size >= c.maxSize {
		c.evict()
	}
	el := c.order.PushFront(&entry{key: key, value: value})
	c.items[key] = el
	c.size++
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.order.Remove(el)
		delete(c.items, key)
		c.size--
	}
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*list.Element)
	c.order.Init()
	c.size = 0
}

func (c *Cache) evict() {
	el := c.order.Back()
	if el == nil {
		return
	}
	c.order.Remove(el)
	kv := el.Value.(*entry)
	delete(c.items, kv.key)
	c.size--
	c.evictions++
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return int(c.size)
}

func (c *Cache) Evictions() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.evictions
}
