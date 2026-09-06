package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type jsPage struct {
	app.Compo
}

func newJSPage() *jsPage { _ = "STUB: not implemented"; return nil }

func (p *jsPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *jsPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *jsPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
