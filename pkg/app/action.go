package app

import (
	"sync"
)

type Action struct {
	Name string

	Value any

	Tags Tags
}

type ActionHandler func(Context, Action)

func Handle(actionName string, h ActionHandler) { _ = "STUB: not implemented"; return }

var actionHandlers = make(map[string]ActionHandler)

type actionHandler struct {
	Source   UI
	Function ActionHandler
	Async    bool
}

type actionManager struct {
	mutex    sync.Mutex
	handlers map[string]map[string]actionHandler
}

func (m *actionManager) Handle(action string, source UI, async bool, handler ActionHandler) {
	_ = "STUB: not implemented"
	return
}

func (m *actionManager) Post(ctx Context, a Action) { _ = "STUB: not implemented"; return }

func (m *actionManager) Cleanup() { _ = "STUB: not implemented"; return }

func actionHandlerKey(source UI, handler ActionHandler) string {
	_ = "STUB: not implemented"
	return ""
}
