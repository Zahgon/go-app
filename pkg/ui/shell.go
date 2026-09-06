package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IShell interface {
	app.UI

	ID(v string) IShell

	Class(v string) IShell

	PaneWidth(px int) IShell

	AdsWidth(px int) IShell

	HamburgerButton(v app.UI) IShell

	HamburgerMenu(v ...app.UI) IShell

	Menu(v ...app.UI) IShell

	Index(v ...app.UI) IShell

	Content(v ...app.UI) IShell

	Ads(v ...app.UI) IShell
}

func Shell() IShell { _ = "STUB: not implemented"; return *new(IShell) }

type shell struct {
	app.Compo

	Iid              string
	Iclass           string
	IpaneWidth       int
	IadsWidth        int
	IhamburgerButton app.UI
	IhamburgerMenu   []app.UI
	Imenu            []app.UI
	Iindex           []app.UI
	Icontent         []app.UI
	Iads             []app.UI

	id                string
	hideMenu          bool
	hideIndex         bool
	hideAds           bool
	showHamburgerMenu bool
	width             int
}

func (s *shell) ID(v string) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) Class(v string) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) PaneWidth(px int) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) AdsWidth(px int) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) HamburgerButton(v app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) HamburgerMenu(v ...app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) Menu(v ...app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) Index(v ...app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) Content(v ...app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) Ads(v ...app.UI) IShell { _ = "STUB: not implemented"; return *new(IShell) }

func (s *shell) OnPreRender(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *shell) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *shell) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *shell) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *shell) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (s *shell) refresh(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *shell) layoutSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *shell) onHamburgerButtonClick(ctx app.Context, e app.Event) {
	_ = "STUB: not implemented"
	return
}

func (s *shell) hideHamburgerMenu(ctx app.Context, e app.Event) { _ = "STUB: not implemented"; return }
