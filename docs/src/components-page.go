package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type componentsPage struct {
	app.Compo
}

func newComponentsPage() *componentsPage { _ = "STUB: not implemented"; return nil }

func (p *componentsPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *componentsPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *componentsPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
