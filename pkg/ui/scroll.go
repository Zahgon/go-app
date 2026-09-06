package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IScroll interface {
	app.UI

	ID(v string) IScroll

	Class(v string) IScroll

	HeaderHeight(px int) IScroll

	Header(v ...app.UI) IScroll

	Content(v ...app.UI) IScroll

	FooterHeight(px int) IScroll

	Footer(v ...app.UI) IScroll
}

func Scroll() IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

type scroll struct {
	app.Compo

	Iid           string
	Iclass        string
	IheaderHeight int
	IfooterHeight int
	Iheader       []app.UI
	Icontent      []app.UI
	Ifooter       []app.UI

	hpadding int
	vpadding int
	width    int
}

func (s *scroll) ID(v string) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) Class(v string) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) HeaderHeight(px int) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) Header(v ...app.UI) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) Content(v ...app.UI) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) FooterHeight(px int) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) Footer(v ...app.UI) IScroll { _ = "STUB: not implemented"; return *new(IScroll) }

func (s *scroll) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *scroll) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *scroll) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *scroll) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (s *scroll) resize(ctx app.Context) { _ = "STUB: not implemented"; return }
