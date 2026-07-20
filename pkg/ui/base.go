package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IBase interface {
	app.UI

	ID(v string) IBase

	Class(v string) IBase

	Content(v ...app.UI) IBase
}

func Base() IBase { _ = "STUB: not implemented"; return *new(IBase) }

type base struct {
	app.Compo

	Iid      string
	Iclass   string
	Icontent []app.UI

	hpadding int
	vpadding int
	width    int
}

func (b *base) ID(v string) IBase { _ = "STUB: not implemented"; return *new(IBase) }

func (b *base) Class(v string) IBase { _ = "STUB: not implemented"; return *new(IBase) }

func (b *base) Content(v ...app.UI) IBase { _ = "STUB: not implemented"; return *new(IBase) }

func (b *base) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *base) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *base) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *base) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (b *base) resize(ctx app.Context) { _ = "STUB: not implemented"; return }
