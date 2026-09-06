package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type concurrencyPage struct {
	app.Compo
}

func newConcurrencyPage() *concurrencyPage { _ = "STUB: not implemented"; return nil }

func (p *concurrencyPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *concurrencyPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *concurrencyPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
