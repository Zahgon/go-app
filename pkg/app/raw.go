package app

func Raw(v string) UI { _ = "STUB: not implemented"; return *new(UI) }

type raw struct {
	jsElement     Value
	parentElement UI
	treeDepth     uint
	tag           string
	value         string
}

func (r *raw) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (r *raw) Mounted() bool { _ = "STUB: not implemented"; return false }

func (r *raw) depth() uint { _ = "STUB: not implemented"; return 0 }

func (r *raw) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (r *raw) setParent(p UI) UI { _ = "STUB: not implemented"; return *new(UI) }

func rawRootTagName(raw string) string { _ = "STUB: not implemented"; return "" }
