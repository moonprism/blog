package cache

import (
	"sync"
	"time"
)

// 不需要redis的分布式功能呢

type CacheItem struct {
	Value      any
	Expiration int64
}

func (item *CacheItem) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	items           map[string]*CacheItem
	mu              sync.RWMutex
	ttl             time.Duration // 默认 TTL（Time To Live）
	cleanupInterval time.Duration // 定时清理间隔
}

func New() *Cache {
	cache := &Cache{
		items:           make(map[string]*CacheItem),
		ttl:             0, // 取消默认过期时间，支持永久缓存
		cleanupInterval: time.Minute,
	}
	cache.StartCleanup()
	return cache
}

func (c *Cache) Set(key string, value any, duration time.Duration) bool {
	// 禁止缓存空值
	if value == nil {
		return false
	}
	var expiration int64
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	} else if c.ttl > 0 {
		expiration = time.Now().Add(c.ttl).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: expiration,
	}
	return true
}

func (c *Cache) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || item.Expired() {
		return nil
	}
	return item.Value
}

func (c *Cache) TTL(key string) time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || item.Expired() {
		return -2
	}
	if item.Expiration == 0 {
		return -1
	}
	ttl := time.Until(time.Unix(0, item.Expiration))
	if ttl < 0 {
		return -2
	}
	return ttl
}

func (c *Cache) Del(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
	return true
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if item.Expired() {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

// 启动任务，清理过期缓存
func (c *Cache) StartCleanup() {
	go c.cleanup()
}
