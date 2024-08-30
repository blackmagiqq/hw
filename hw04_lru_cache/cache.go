package hw04lrucache

import (
	"maps"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	valueFromQueue, keyIsInQueue := c.items[key]

	if keyIsInQueue {
		c.queue.MoveToFront(valueFromQueue)
		c.items[key].Value = value
	} else {
		valueFromQueue = c.queue.PushFront(value)
		c.items[key] = valueFromQueue
	}

	if len(c.items) > c.capacity {
		backEl := c.queue.Back()
		c.queue.Remove(backEl)
		maps.DeleteFunc(c.items, func(_ Key, value *ListItem) bool {
			return value == backEl
		})
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
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
