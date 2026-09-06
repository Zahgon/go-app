package cache

import (
	"context"
	"reflect"
)

type Cache interface {
	Get(ctx context.Context, key string) (Item, bool)

	Set(ctx context.Context, key string, i Item)

	Del(ctx context.Context, key string)

	Len() int

	Size() int
}

type Item interface {
	Size() int
}

type Bytes []byte

func (b Bytes) Size() int { _ = "STUB: not implemented"; return 0 }

type String string

func (s String) Size() int { _ = "STUB: not implemented"; return 0 }

type Int int

func (i Int) Size() int { _ = "STUB: not implemented"; return 0 }

type Float float64

func (f Float) Size() int { _ = "STUB: not implemented"; return 0 }

var (
	intSize   = int(reflect.TypeOf(42).Size())
	floatSize = int(reflect.TypeOf(23.42).Size())
)
