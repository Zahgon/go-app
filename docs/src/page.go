package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

const (
	headerHeight  = 72
	adsenseClient = "ca-pub-1013306768105236"
	adsenseSlot   = "9307554044"
)

type page struct {
	app.Compo

	Iclass   string
	Iindex   []app.UI
	Iicon    string
	Ititle   string
	Icontent []app.UI

	updateAvailable bool
}

func newPage() *page { _ = "STUB: not implemented"; return nil }

func (p *page) Index(v ...app.UI) *page { _ = "STUB: not implemented"; return nil }

func (p *page) Icon(v string) *page { _ = "STUB: not implemented"; return nil }

func (p *page) Title(v string) *page { _ = "STUB: not implemented"; return nil }

func (p *page) Content(v ...app.UI) *page { _ = "STUB: not implemented"; return nil }

func (p *page) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *page) OnAppUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *page) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (p *page) updateApp(ctx app.Context, e app.Event) { _ = "STUB: not implemented"; return }

func scrollTo(ctx app.Context) { _ = "STUB: not implemented"; return }

func fragmentFocus(fragment string) string { _ = "STUB: not implemented"; return "" }
