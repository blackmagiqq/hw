package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity   int
	queue      List
	items      map[Key]*ListItem
	valueToKey map[*ListItem]Key
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	valueFromQueue, keyIsInQueue := c.items[key]

	if keyIsInQueue {
		c.queue.MoveToFront(valueFromQueue)
		c.items[key].Value = value
		return keyIsInQueue
	}

	valueFromQueue = c.queue.PushFront(value)
	c.items[key] = valueFromQueue
	c.valueToKey[valueFromQueue] = key

	if len(c.items) > c.capacity {
		backEl := c.queue.Back()
		c.queue.Remove(backEl)

		key := c.valueToKey[backEl]

		delete(c.items, key)
		delete(c.valueToKey, backEl)
	}

	return keyIsInQueue
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	value, keyIsInQueue := c.items[key]
	if !keyIsInQueue {
		return nil, false
	}
	c.queue.MoveToFront(value)
	return value.Value, true
}

func (c *lruCache) Clear() {
	for k, v := range c.items {
		c.queue.Remove(v)
		delete(c.items, k)
		delete(c.valueToKey, v)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity:   capacity,
		queue:      NewList(),
		items:      make(map[Key]*ListItem, capacity),
		valueToKey: make(map[*ListItem]Key, capacity),
	}
}
