package src

import (
	"GeeCache/src/lru"
	"sync"
)

/*
使用mutex 封装 Cache 的add 和 get 方法，
*/
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 缓存未初始化
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}

func (c *cache) delete(key string) (ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ok := c.lru.Remove(key); ok {
		return ok
	}
	return
}
