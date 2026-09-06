package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type architecturePage struct {
	app.Compo
}

func newArchitecturePage() *architecturePage { _ = "STUB: not implemented"; return nil }

func (p *architecturePage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *architecturePage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *architecturePage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
