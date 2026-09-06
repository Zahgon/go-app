package app

type Tagger interface {
	Tags() Tags
}

type Tags map[string]string

func (t Tags) Tags() Tags { _ = "STUB: not implemented"; return *new(Tags) }

func (t Tags) Set(name string, v any) { _ = "STUB: not implemented"; return }

func (t Tags) Get(name string) string { _ = "STUB: not implemented"; return "" }

type Tag struct {
	Name  string
	Value string
}

func (t Tag) Tags() Tags { _ = "STUB: not implemented"; return *new(Tags) }

func T(name string, value any) Tag { _ = "STUB: not implemented"; return *new(Tag) }
