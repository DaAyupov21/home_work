package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type entry struct {
	key   Key
	value interface{}
}

type lruCache struct {
	cap   int
	q     List
	items map[Key]*ListItem

	mu sync.Mutex // делаем потокобезопасным по умолчанию (звёздочка)
}

func NewCache(capacity int) Cache {
	if capacity < 1 {
		capacity = 1
	}
	return &lruCache{
		cap:   capacity,
		q:     NewList(),
		items: make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if it, ok := c.items[key]; ok {
		it.Value.(*entry).value = value
		c.q.MoveToFront(it)
		return true
	}

	it := c.q.PushFront(&entry{key: key, value: value})
	c.items[key] = it

	if c.q.Len() > c.cap {
		tail := c.q.Back()
		if tail != nil {
			old := tail.Value.(*entry)
			delete(c.items, old.key)
			c.q.Remove(tail)
		}
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if it, ok := c.items[key]; ok {
		c.q.MoveToFront(it)
		return it.Value.(*entry).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.q = NewList()
	c.items = make(map[Key]*ListItem, c.cap)
}
