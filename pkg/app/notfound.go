package app

var (
	NotFound UI = &notFound{}
)

type notFound struct {
	Compo
	Icon string
}

func (n *notFound) OnMount(Context) { _ = "STUB: not implemented"; return }

func (n *notFound) Render() UI { _ = "STUB: not implemented"; return *new(UI) }
