package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("aaa", 400) // [400]
		c.Set("bbb", 300) // [300, 400]
		c.Set("ccc", 200) // [200, 300, 400]
		c.Set("ddd", 100) // [100, 200, 300]

		cache := c.(*lruCache)
		list := cache.queue

		expectedFrontValue := 100
		expectedBackValue := 300

		require.Equal(t, expectedFrontValue, list.Front().Value, "unexpected front value")
		require.Equal(t, expectedBackValue, list.Back().Value, "unexpected back value")

		c.Set("bbb", 3000) // [3000, 100, 200]
		c.Set("ddd", 1000) // [1000, 3000, 200]
		c.Set("fff", 4000) // [4000, 1000, 3000]

		expectedFrontValue = 4000
		expectedBackValue = 3000

		require.Equal(t, expectedFrontValue, list.Front().Value, "unexpected front value")
		require.Equal(t, expectedBackValue, list.Back().Value, "unexpected back value")
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
