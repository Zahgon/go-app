package app

import (
	"encoding/json"
	"sync"
	"time"
)

type State struct {
	value     any
	expiresAt time.Time

	ctx       Context
	name      string
	expire    func(State, time.Time) State
	persist   func(State, bool) State
	broadcast func(State) State
}

func (s State) ExpiresIn(v time.Duration) State { _ = "STUB: not implemented"; return *new(State) }

func (s State) ExpiresAt(v time.Time) State { _ = "STUB: not implemented"; return *new(State) }

func (s State) Persist() State { _ = "STUB: not implemented"; return *new(State) }

func (s State) PersistWithEncryption() State { _ = "STUB: not implemented"; return *new(State) }

func (s State) Broadcast() State { _ = "STUB: not implemented"; return *new(State) }

type storableState struct {
	Value          json.RawMessage `json:",omitempty"`
	EncryptedValue []byte          `json:",omitempty"`
	ExpiresAt      time.Time       `json:",omitempty"`
}

type Observer struct {
	source        UI
	receiver      any
	condition     func() bool
	changeHandler func()
	broadcast     bool

	state           string
	setObserver     func(Observer) Observer
	enableBroadcast func()
}

func (o Observer) While(condition func() bool) Observer {
	_ = "STUB: not implemented"
	return *new(Observer)
}

func (o Observer) OnChange(h func()) Observer { _ = "STUB: not implemented"; return *new(Observer) }

func (o Observer) WithBroadcast() Observer { _ = "STUB: not implemented"; return *new(Observer) }

func (o Observer) observing() bool { _ = "STUB: not implemented"; return false }

type stateManager struct {
	mutex             sync.RWMutex
	states            map[string]State
	observers         map[string]map[UI]Observer
	initBroadcastOnce sync.Once
	broadcastStoreID  string
	broadcastChannel  Value
}

func (m *stateManager) Observe(ctx Context, state string, receiver any) Observer {
	_ = "STUB: not implemented"
	return *new(Observer)
}

func (m *stateManager) setObserver(v Observer) Observer {
	_ = "STUB: not implemented"
	return *new(Observer)
}

func (m *stateManager) Get(ctx Context, state string, receiver any) {
	_ = "STUB: not implemented"
	return
}

func (m *stateManager) getStoredState(ctx Context, state string, receiver any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *stateManager) Set(ctx Context, state string, v any) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (m *stateManager) setExpiration(s State, v time.Time) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (m *stateManager) persist(s State, encrypt bool) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (m *stateManager) broadcast(s State) State { _ = "STUB: not implemented"; return *new(State) }

func (m *stateManager) initBroadcast(ctx Context) { _ = "STUB: not implemented"; return }

func (m *stateManager) handleBroadcast(ctx Context, data Value) { _ = "STUB: not implemented"; return }

func (m *stateManager) Delete(ctx Context, state string) { _ = "STUB: not implemented"; return }

func (m *stateManager) Cleanup() { _ = "STUB: not implemented"; return }

func (m *stateManager) CleanupExpiredPersistedStates(ctx Context) {
	_ = "STUB: not implemented"
	return
}

func storeValue(recv, v any) error { _ = "STUB: not implemented"; return nil }

func expiredTime(v time.Time) bool { _ = "STUB: not implemented"; return false }
