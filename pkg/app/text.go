package app

func Text(v any) UI { _ = "STUB: not implemented"; return *new(UI) }

func Textf(format string, v ...any) UI { _ = "STUB: not implemented"; return *new(UI) }

type text struct {
	jsvalue       Value
	parentElement UI
	value         string
}

func (t *text) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (t *text) Mounted() bool { _ = "STUB: not implemented"; return false }

func (t *text) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (t *text) setParent(p UI) UI { _ = "STUB: not implemented"; return *new(UI) }
