package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IIcon interface {
	app.UI

	ID(v string) IIcon

	Class(v string) IIcon

	Style(k, v string) IIcon

	Size(px int) IIcon

	Src(v string) IIcon
}

func Icon() IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

type icon struct {
	app.Compo

	Iid     string
	Iclass  string
	Istyles []style
	Isize   int
	Isrc    string
}

func (i *icon) ID(v string) IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

func (i *icon) Class(v string) IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

func (i *icon) Style(k, v string) IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

func (i *icon) Size(px int) IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

func (i *icon) Src(v string) IIcon { _ = "STUB: not implemented"; return *new(IIcon) }

func (i *icon) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func isSVG(v string) bool { _ = "STUB: not implemented"; return false }
