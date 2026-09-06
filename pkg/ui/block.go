package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IBlock interface {
	app.UI

	ID(v string) IBlock

	Class(v string) IBlock

	Top() IBlock

	Middle() IBlock

	Padding(v bool) IBlock

	MaxContentWidth(px int) IBlock

	Content(v ...app.UI) IBlock
}

func Block() IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

type block struct {
	app.Compo

	Iid              string
	Iclass           string
	Ialignment       alignment
	ImaxContentWidth int
	Ipadding         bool
	Icontent         []app.UI

	padding int
	width   int
}

func (b *block) ID(v string) IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) Class(v string) IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) Top() IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) Middle() IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) MaxContentWidth(px int) IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) Padding(v bool) IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) Content(v ...app.UI) IBlock { _ = "STUB: not implemented"; return *new(IBlock) }

func (b *block) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *block) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *block) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (b *block) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (b *block) resize(ctx app.Context) { _ = "STUB: not implemented"; return }
