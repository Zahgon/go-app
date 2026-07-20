package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IFlow interface {
	app.UI

	ID(v string) IFlow

	Class(v string) IFlow

	ItemWidth(px int) IFlow

	Spacing(px int) IFlow

	StretchItems() IFlow

	Content(elems ...app.UI) IFlow
}

func Flow() IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

type flow struct {
	app.Compo

	Iid           string
	Iclass        string
	IitemWidth    int
	Ispacing      int
	IstretchItems bool
	Icontent      []app.UI

	id          string
	itemsPerRow int
	itemWidth   float64
}

func (f *flow) ID(v string) IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) Class(v string) IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) ItemWidth(px int) IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) Spacing(px int) IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) StretchItems() IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) Content(elems ...app.UI) IFlow { _ = "STUB: not implemented"; return *new(IFlow) }

func (f *flow) OnPreRender(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flow) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flow) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flow) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flow) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (f *flow) refresh(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flow) layoutSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }
