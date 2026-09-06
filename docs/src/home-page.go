package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage { _ = "STUB: not implemented"; return nil }

func (p *homePage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *homePage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *homePage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
