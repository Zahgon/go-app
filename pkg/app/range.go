package app

type RangeLoop interface {
	UI

	Slice(f func(int) UI) RangeLoop

	Map(f func(string) UI) RangeLoop

	body() []UI
}

func Range(src any) RangeLoop { _ = "STUB: not implemented"; return *new(RangeLoop) }

type rangeLoop struct {
	children []UI
	source   any
}

func (r rangeLoop) Slice(f func(int) UI) RangeLoop {
	_ = "STUB: not implemented"
	return *new(RangeLoop)
}

func (r rangeLoop) Map(f func(string) UI) RangeLoop {
	_ = "STUB: not implemented"
	return *new(RangeLoop)
}

func (r rangeLoop) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (r rangeLoop) Mounted() bool { _ = "STUB: not implemented"; return false }

func (r rangeLoop) setParent(UI) UI { _ = "STUB: not implemented"; return *new(UI) }

func (r rangeLoop) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (r rangeLoop) body() []UI { _ = "STUB: not implemented"; return nil }
