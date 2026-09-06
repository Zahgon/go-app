package app

type EventHandler func(ctx Context, e Event)

type Event struct {
	Value
}

func (e Event) PreventDefault() { _ = "STUB: not implemented"; return }

func (e Event) StopImmediatePropagation() { _ = "STUB: not implemented"; return }

type EventOption struct {
	name  string
	value string
}

func EventScope(v ...any) EventOption { _ = "STUB: not implemented"; return *new(EventOption) }

func PassiveEvent() EventOption { _ = "STUB: not implemented"; return *new(EventOption) }

type eventHandlers map[string]eventHandler

func (h eventHandlers) Set(event string, eh EventHandler, options ...EventOption) {
	_ = "STUB: not implemented"
	return
}

type eventHandler struct {
	event     string
	scope     string
	passive   bool
	goHandler EventHandler
	jsHandler Func
	close     func()
}

func makeEventHandler(event string, h EventHandler, options ...EventOption) eventHandler {
	_ = "STUB: not implemented"
	return *new(eventHandler)
}

func (h eventHandler) Equal(v eventHandler) bool { _ = "STUB: not implemented"; return false }

func (h eventHandler) options() map[string]any { _ = "STUB: not implemented"; return nil }

func trackMousePosition(e Event) { _ = "STUB: not implemented"; return }
