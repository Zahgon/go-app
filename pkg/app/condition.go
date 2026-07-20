package app

type Condition interface {
	UI

	ElseIf(expr bool, elem func() UI) Condition

	ElseIfSlice(expr bool, elems func() []UI) Condition

	Else(elem func() UI) Condition

	ElseSlice(elems func() []UI) Condition

	body() []UI
}

func If(expr bool, elem func() UI) Condition { _ = "STUB: not implemented"; return *new(Condition) }

func IfSlice(expr bool, elems func() []UI) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

type condition struct {
	children []UI
	matched  bool
}

func (c condition) ElseIf(expr bool, elem func() UI) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

func (c condition) ElseIfSlice(expr bool, elems func() []UI) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

func (c condition) Else(elem func() UI) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

func (c condition) ElseSlice(elems func() []UI) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

func (c condition) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (c condition) Mounted() bool { _ = "STUB: not implemented"; return false }

func (c condition) body() []UI { _ = "STUB: not implemented"; return nil }

func (c condition) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (c condition) setParent(UI) UI { _ = "STUB: not implemented"; return *new(UI) }
