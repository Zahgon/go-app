package app

import (
	"sync"
)

type BrowserStorage interface {
	Set(k string, v any) error

	Get(k string, v any) error

	Del(k string)

	Len() int

	Clear()

	ForEach(f func(k string))

	Contains(k string) bool
}

type memoryStorage struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func newMemoryStorage() *memoryStorage { _ = "STUB: not implemented"; return nil }

func (s *memoryStorage) Set(k string, v any) error { _ = "STUB: not implemented"; return nil }

func (s *memoryStorage) Get(k string, v any) error { _ = "STUB: not implemented"; return nil }

func (s *memoryStorage) Del(k string) { _ = "STUB: not implemented"; return }

func (s *memoryStorage) Clear() { _ = "STUB: not implemented"; return }

func (s *memoryStorage) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *memoryStorage) ForEach(f func(key string)) { _ = "STUB: not implemented"; return }

func (s *memoryStorage) Contains(k string) bool { _ = "STUB: not implemented"; return false }

type jsStorage struct {
	name  string
	mutex sync.RWMutex
}

func newJSStorage(name string) *jsStorage { _ = "STUB: not implemented"; return nil }

func (s *jsStorage) Set(k string, v any) (err error) { _ = "STUB: not implemented"; return nil }

func (s *jsStorage) Get(k string, v any) error { _ = "STUB: not implemented"; return nil }

func (s *jsStorage) Del(k string) { _ = "STUB: not implemented"; return }

func (s *jsStorage) Clear() { _ = "STUB: not implemented"; return }

func (s *jsStorage) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *jsStorage) len() int { _ = "STUB: not implemented"; return 0 }

func (s *jsStorage) ForEach(f func(key string)) { _ = "STUB: not implemented"; return }

func (s *jsStorage) Contains(k string) bool { _ = "STUB: not implemented"; return false }
