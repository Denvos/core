package eviction

import "container/list"

type Policy interface {
	OnAccess(key string)
	OnInsert(key string)
	OnDelete(key string)
	Evict() string
}

type LRU struct {
	order *list.List
	items map[string]*list.Element
}

func NewLRU() *LRU {
	return &LRU{
		order: list.New(),
		items: make(map[string]*list.Element),
	}
}

func (l *LRU) OnAccess(key string) {
	if el, ok := l.items[key]; ok {
		l.order.MoveToFront(el)
	}
}

func (l *LRU) OnInsert(key string) {
	if _, ok := l.items[key]; ok {
		l.order.Remove(l.items[key])
	}
	el := l.order.PushFront(key)
	l.items[key] = el
}

func (l *LRU) OnDelete(key string) {
	if el, ok := l.items[key]; ok {
		l.order.Remove(el)
		delete(l.items, key)
	}
}

func (l *LRU) Evict() string {
	el := l.order.Back()
	if el == nil {
		return ""
	}
	key := el.Value.(string)
	l.order.Remove(el)
	delete(l.items, key)
	return key
}

type LFU struct {
	// implement LFU
}

type FIFO struct {
	order *list.List
}

func NewFIFO() *FIFO {
	return &FIFO{order: list.New()}
}
