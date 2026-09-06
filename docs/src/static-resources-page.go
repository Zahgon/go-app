package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type staticResourcesPage struct {
	app.Compo
}

func newStaticResourcePage() *staticResourcesPage { _ = "STUB: not implemented"; return nil }

func (p *staticResourcesPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *staticResourcesPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *staticResourcesPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
