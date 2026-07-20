package app

import (
	"sync"
)

type cacheItem struct {
	Path string

	ContentType string

	ContentEncoding string

	Body []byte
}

func (i cacheItem) Len() int { _ = "STUB: not implemented"; return 0 }

type memoryCache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
}

func newMemoryCache(size int) *memoryCache { _ = "STUB: not implemented"; return nil }

func (c *memoryCache) Set(i cacheItem) { _ = "STUB: not implemented"; return }

func (c *memoryCache) Get(path string) (cacheItem, bool) {
	_ = "STUB: not implemented"
	return *new(cacheItem), false
}
