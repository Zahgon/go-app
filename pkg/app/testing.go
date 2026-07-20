package app

type TestEngine interface {
	Load(Composer) error

	ConsumeNext()

	ConsumeAll()
}

func NewTestEngine() TestEngine { _ = "STUB: not implemented"; return *new(TestEngine) }

func Match(expected UI, root UI, path ...int) error { _ = "STUB: not implemented"; return nil }

type TestUIDescriptor struct {
	Path []int

	Expected UI
}

func TestPath(p ...int) []int { _ = "STUB: not implemented"; return nil }

func TestMatch(root UI, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }

func match(n UI, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }

func matchText(n *text, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }

func matchHTML(n HTML, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }

func matchHTMLAttributes(a, b attributes) error { _ = "STUB: not implemented"; return nil }

func matchHTMLEventHandlers(a, b eventHandlers) error { _ = "STUB: not implemented"; return nil }

func matchComponent(n Composer, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }

func matchRaw(n *raw, d TestUIDescriptor) error { _ = "STUB: not implemented"; return nil }
