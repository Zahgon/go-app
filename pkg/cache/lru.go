package cache

import (
	"context"
	"sync"
	"time"
)

type LRU struct {
	MaxSize int

	ItemTTL time.Duration

	OnEvict func(key string, i Item)

	once     sync.Once
	mutex    sync.Mutex
	size     int
	items    map[string]*lruItem
	priority []*lruItem
}

func (c *LRU) Get(ctx context.Context, key string) (Item, bool) {
	_ = "STUB: not implemented"
	return *new(Item), false
}

func (c *LRU) Set(ctx context.Context, key string, i Item) { _ = "STUB: not implemented"; return }

func (c *LRU) Del(ctx context.Context, key string) { _ = "STUB: not implemented"; return }

func (c *LRU) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *LRU) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *LRU) init() {
	if c.MaxSize <= 0 {
		c.MaxSize = 16000000
	}

	c.items = make(map[string]*lruItem, 64)
	c.priority = make([]*lruItem, 0, len(c.items))
}

func (c *LRU) free(size int) { _ = "STUB: not implemented"; return }

func (c *LRU) removeLastItem() { _ = "STUB: not implemented"; return }

func (c *LRU) add(i *lruItem) { _ = "STUB: not implemented"; return }

type lruItem struct {
	key       string
	count     int
	expiresAt time.Time
	value     Item
}

func (i *lruItem) priority(now time.Time) int { _ = "STUB: not implemented"; return 0 }

func (i *lruItem) IsExpired(now time.Time) bool { _ = "STUB: not implemented"; return false }

func sortLRUItems(now time.Time, v []*lruItem) { _ = "STUB: not implemented"; return }
