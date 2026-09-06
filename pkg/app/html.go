package app

type HTML interface {
	UI

	Tag() string

	XMLNamespace() string

	SelfClosing() bool

	depth() uint
	attrs() attributes
	setAttrs(attributes) HTML
	events() eventHandlers
	setEvents(eventHandlers) HTML
	setDepth(uint) UI
	setJSElement(Value) HTML
	parent() UI
	body() []UI
	setBody([]UI) HTML
}

type htmlElement struct {
	tag           string
	xmlns         string
	treeDepth     uint
	isSelfClosing bool
	jsElement     Value
	attributes    attributes
	eventHandlers eventHandlers
	parentElement UI
	children      []UI
}

func (e *htmlElement) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (e *htmlElement) Mounted() bool { _ = "STUB: not implemented"; return false }

func (e *htmlElement) Tag() string { _ = "STUB: not implemented"; return "" }

func (e *htmlElement) XMLNamespace() string { _ = "STUB: not implemented"; return "" }

func (e *htmlElement) SelfClosing() bool { _ = "STUB: not implemented"; return false }

func (e *htmlElement) depth() uint { _ = "STUB: not implemented"; return 0 }

func (e *htmlElement) attrs() attributes { _ = "STUB: not implemented"; return *new(attributes) }

func (e *htmlElement) setAttr(name string, value any) { _ = "STUB: not implemented"; return }

func (e *htmlElement) events() eventHandlers { _ = "STUB: not implemented"; return *new(eventHandlers) }

func (e *htmlElement) setEventHandler(event string, h EventHandler, options ...EventOption) {
	_ = "STUB: not implemented"
	return
}

func (e *htmlElement) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (e *htmlElement) body() []UI { _ = "STUB: not implemented"; return nil }
