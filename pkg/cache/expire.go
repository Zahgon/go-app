package cache

import (
	"context"
	"sync"
	"time"
)

type Expire struct {
	ItemTTL time.Duration

	once  sync.Once
	mutex sync.RWMutex
	size  int
	items map[string]*memItem
	queue []*memItem
}

func (c *Expire) Get(ctx context.Context, key string) (Item, bool) {
	_ = "STUB: not implemented"
	return *new(Item), false
}

func (c *Expire) Set(ctx context.Context, key string, i Item) { _ = "STUB: not implemented"; return }

func (c *Expire) Del(ctx context.Context, key string) { _ = "STUB: not implemented"; return }

func (c *Expire) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Expire) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *Expire) init() {
	c.items = make(map[string]*memItem)
}

func (c *Expire) add(i *memItem) { _ = "STUB: not implemented"; return }

func (c *Expire) del(i *memItem) { _ = "STUB: not implemented"; return }

func (c *Expire) expire() { _ = "STUB: not implemented"; return }

type memItem struct {
	key       string
	expiresAt time.Time
	value     Item
}

func (i *memItem) isExpired(now time.Time) bool { _ = "STUB: not implemented"; return false }
